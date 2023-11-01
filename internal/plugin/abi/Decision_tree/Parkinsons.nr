// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;22],) -> pub u3 {
	if inputs[21] <= 45 {
		if inputs[7] <= 2 {
			1
		} else {
			if inputs[21] <= 33 {
				0
			} else {
				if inputs[21] <= 38 {
					if inputs[10] <= 41 {
						1
					} else {
						0
					}
				} else {
					0
				}
			}
		}
	} else {
		if inputs[15] <= 172 {
			if inputs[18] <= 67 {
				0
			} else {
				if inputs[19] <= 93 {
					if inputs[17] <= 154 {
						0
					} else {
						1
					}
				} else {
					if inputs[16] <= 241 {
						1
					} else {
						if inputs[16] <= 243 {
							0
						} else {
							1
						}
					}
				}
			}
		} else {
			if inputs[16] <= 100 {
				if inputs[6] <= 12 {
					1
				} else {
					0
				}
			} else {
				if inputs[6] <= 8 {
					0
				} else {
					1
				}
			}
		}
	}
}