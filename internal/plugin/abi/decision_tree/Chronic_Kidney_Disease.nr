// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;24],) -> pub u3 {
	if inputs[14] <= 168 {
		1
	} else {
		if inputs[4] <= 26 {
			0
		} else {
			1
		}
	}
}
