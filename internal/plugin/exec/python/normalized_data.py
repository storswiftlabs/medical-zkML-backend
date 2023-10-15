import argparse
import os


def str_to_float(data):
    nums = []
    for i in range(len(data)):
        nums.append(float(data[i]))
    return nums


parser = argparse.ArgumentParser()

# 添加参数
parser.add_argument('params', nargs='+', help='参数列表')

# 解析参数
args = parser.parse_args()

# 访问参数数组
disease = args.params[0]
disease = disease.replace(" ", "_")
params = args.params[1:]
params = str_to_float(params)
# print(params)

with open(f'{os.getcwd()}/internal/plugin/exec/python/column_uniformization_info/{disease}.tsv', 'r') as f:
    lines = f.readlines()


max = str_to_float(lines[0].strip().split('\t')[1:])
min = str_to_float(lines[1].strip().split('\t')[1:])


for index in range(0, len(params)):
    if params[index] > max[index]:
        params[index] = max[index]

# print(max)
# print(min)

result = []

for index in range(0, len(params)):
    # 小于0， 设置为0， 大于1， 设置为1
    molecule = params[index] - min[index]
    if molecule == 0:
        result.append(0)
        continue
    denominator = max[index] - min[index]
    result.append((params[index] - min[index]) / (max[index] - min[index]))

for index in range(0, len(result)):
    if result[index] < 0: result[index] = 0
    if result[index] > 1: result[index] = 1

print(result)
