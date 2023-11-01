// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;18],) -> pub u3 {
	if inputs[17] <= 55 {
		if inputs[1] <= 128 {
			if inputs[5] <= 128 {
				if inputs[12] <= 213 {
					if inputs[17] <= 18 {
						if inputs[14] <= 64 {
							1
						} else {
							if inputs[15] <= 128 {
								1
							} else {
								2
							}
						}
					} else {
						if inputs[13] <= 128 {
							if inputs[14] <= 191 {
								1
							} else {
								if inputs[15] <= 128 {
									1
								} else {
									2
								}
							}
						} else {
							if inputs[9] <= 43 {
								1
							} else {
								2
							}
						}
					}
				} else {
					2
				}
			} else {
				if inputs[12] <= 43 {
					0
				} else {
					2
				}
			}
		} else {
			if inputs[12] <= 213 {
				1
			} else {
				if inputs[14] <= 191 {
					1
				} else {
					2
				}
			}
		}
	} else {
		if inputs[14] <= 191 {
			if inputs[17] <= 91 {
				1
			} else {
				if inputs[5] <= 128 {
					1
				} else {
					2
				}
			}
		} else {
			if inputs[8] <= 64 {
				if inputs[4] <= 128 {
					if inputs[15] <= 128 {
						if inputs[10] <= 191 {
							2
						} else {
							1
						}
					} else {
						2
					}
				} else {
					if inputs[13] <= 128 {
						1
					} else {
						2
					}
				}
			} else {
				3
			}
		}
	}
}
