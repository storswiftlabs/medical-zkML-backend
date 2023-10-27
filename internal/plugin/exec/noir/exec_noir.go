package noir

import (
	"bufio"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/db"
	m "medical-zkml-backend/internal/module"
	"medical-zkml-backend/pkg/config"
	"medical-zkml-backend/pkg/utils"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"
)

var rootPath string
var msg = "Noir disease prediction"

func init() {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(path.Dir(path.Dir(filename)))))

}

type ExecNoir struct {
	root  string   // Command source file path
	dir   string   // Project Root Path
	tasks []string // task list
}

var Noir *ExecNoir

func InitNoir(path, dir string) *ExecNoir {
	Noir = &ExecNoir{
		root: path,
		dir:  dir,
	}
	return Noir
}

func (n *ExecNoir) NoirVersionCheck() {
	cmd := exec.Command(n.root, "--version")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	if strings.Split(strings.TrimSpace(string(output)), " ")[1] != config.GetConfig().GetString("nargo.version") {
		fmt.Printf("Please use nargo version 0.10.5")
		os.Exit(1)
	}
}

func (n *ExecNoir) PathCheck() {
	circuitPath := fmt.Sprintf("%s/%s", rootPath, n.dir)
	_, err := os.Stat(circuitPath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(circuitPath, os.ModePerm); err != nil {
			fmt.Printf("Could not create circuit: %s", err)
			os.Exit(2)
		}
	}
}

func (n *ExecNoir) prediction(path, disease, module string, quantized *m.Quantized) error {

	if err := n.circuitInjection(path, disease, module); err != nil {
		return err
	}

	if err := n.quantitativeInjection(path); err != nil {
		return err
	}

	// Information injection
	if err := n.inputsToProver(quantized.Data, path, quantized.Scale, quantized.ZeroPoint); err != nil {
		return err
	}

	// Project execution
	cmd := exec.Command(n.root, "prove", "p")
	cmd.Dir = fmt.Sprintf("%s/internal/plugin/circuit/%s", rootPath, path)
	if err := cmd.Run(); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", path))))
		return err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s execution prediction", path), "success"))
	return nil
}

func (n *ExecNoir) DiseasePrediction(user string, recordPrediction *m.RecordPrediction, disease, module string, quantized *m.Quantized) {

	zap.L().Info(msg, zap.String("Status", "start"))

	module = strings.ReplaceAll(strings.ReplaceAll(module, " ", "_"), "-", "_")
	disease = strings.ReplaceAll(strings.ReplaceAll(disease, " ", "_"), "-", "_")
	project, err := n.createProject(recordPrediction.ID, disease, module)
	if err != nil {
		// Update the status of the database
		recordPrediction.Status = m.CreateFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Failed to create noir project", err.Error())
		_ = db.RecordPredict(recordPrediction)
		return
	}
	// Successfully created, updating database status
	recordPrediction.Status = m.CreateSuccess
	_ = db.RecordPredict(recordPrediction)

	//
	if err := n.prediction(project, disease, module, quantized); err != nil {
		// Update database status
		recordPrediction.Status = m.PredictionFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Noir circuit failed to perform prediction", err.Error())
		_ = db.RecordPredict(recordPrediction)
		return
	}
	// Execution completed, update database status
	recordPrediction.Status = m.PredictionComplete
	_ = db.RecordPredict(recordPrediction)

	// Obtain results and update database status
	result, err := n.getResult(project)
	if err != nil {
		recordPrediction.Status = m.GetResultFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Failed to obtain prediction results", err.Error())
		_ = db.RecordPredict(recordPrediction)
		return
	}
	result = strings.Trim(strings.TrimSpace(strings.Split(result, "= ")[1]), "\"")

	proof, err := n.getProof(project)
	if err != nil {
		recordPrediction.Status = m.GetResultFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Failed to obtain prediction results", err.Error())
		_ = db.RecordPredict(recordPrediction)
		return
	}
	proof = "0x" + proof
	// Store prediction information
	// TODO: Store prediction results and modify method names
	_ = db.SavePredictedResult(user, disease, module, result, proof, recordPrediction.ID)

	// end
	recordPrediction.Status = m.Complete
	recordPrediction.Output = utils.Hexadecimal2Decimal(result)
	recordPrediction.EndTime = time.Now().Unix()
	_ = db.RecordPredict(recordPrediction)

	_ = n.cleanProject(project)

	return

}

func (n *ExecNoir) getResult(project string) (string, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s/%s", n.dir, project, "Verifier.toml"))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result string
	if scanner.Scan() {
		result = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return "", err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s get result", project), "success"))
	return result, nil
}

func (n *ExecNoir) cleanProject(project string) error {
	if err := os.RemoveAll(fmt.Sprintf("%s/internal/plugin/circuit/%s", rootPath, project)); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s clean project", project), "success"))
	return nil
}

func (n *ExecNoir) getProof(project string) (string, error) {

	file, err := os.Open(fmt.Sprintf("%s/%s/proofs/p.proof", n.dir, project))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var proof string
	if scanner.Scan() {
		proof = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s get proof", project), "success"))
	return proof, nil
}

func (n *ExecNoir) inputsToProver(inputs []string, project, scale, zeroPoint string) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/%s", rootPath, project, "Prover.toml"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, fmt.Sprintf("inputs = [%v]", utils.StringList2String(inputs)))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	_, err = fmt.Fprintln(file, fmt.Sprintf("inputs_scale = %s", scale))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	_, err = fmt.Fprintln(file, fmt.Sprintf("inputs_zero_point = %s", zeroPoint))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", project))))
		return err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s inputs injection", project), "success"))
	return nil
}

func (n *ExecNoir) createProject(id uint, disease, module string) (string, error) {
	// nargo new module_id

	newProject := fmt.Sprintf("%s_%s_%d", module, disease, id)
	cmd := exec.Command(n.root, "new", newProject)
	cmd.Dir = n.dir
	if err := cmd.Run(); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", newProject))))
		return "", errors.New(fmt.Sprintf("Failed to create task: %s", err.Error()))
	}
	zap.L().Info(msg, zap.String("create newProject success", newProject))
	return newProject, nil
}

func (n *ExecNoir) circuitInjection(path, disease, module string) error {
	content, err := os.ReadFile(fmt.Sprintf("%s/internal/plugin/abi/%s/%s.nr", rootPath, module, disease))
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", path))))
		return err
	}

	if err := os.WriteFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/src/main.nr", rootPath, path), content, 0644); err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", path))))
		return err
	}
	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s circuit injection", path), "success"))
	return nil
}

func (n *ExecNoir) quantitativeInjection(path string) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/internal/plugin/circuit/%s/Nargo.toml", rootPath, path), os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		zap.L().Error(msg, zap.Error(err))
		zap.L().Error(msg, zap.Error(errors.New(fmt.Sprintf("Task %s execution failed", path))))
		return err
	}
	defer file.Close()

	newLine := "\nquantization_arithmetic = { path = \"../../abi/quantization_arithmetic\"}\n"

	writer := bufio.NewWriter(file)
	writer.WriteString(newLine)
	writer.Flush()

	zap.L().Info(msg, zap.String(fmt.Sprintf("Task %s quantitative injection", path), "success"))
	return nil
}
