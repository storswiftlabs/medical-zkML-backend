// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;6],) -> pub u3 {
	if inputs[2] <= 128 {
		if inputs[3] <= 128 {
			0
		} else {
			1
		}
	} else {
		if inputs[0] <= 104 {
			0
		} else {
			if inputs[4] <= 128 {
				2
			} else {
				if inputs[3] <= 128 {
					2
				} else {
					3
				}
			}
		}
	}
}