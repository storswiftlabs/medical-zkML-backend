// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;12],) -> pub u3 {
	if inputs[11] <= 58 {
		if inputs[4] <= 226 {
			if inputs[2] <= 4 {
				if inputs[11] <= 22 {
					if inputs[4] <= 52 {
						if inputs[4] <= 33 {
							1
						} else {
							0
						}
					} else {
						1
					}
				} else {
					if inputs[6] <= 93 {
						if inputs[6] <= 50 {
							1
						} else {
							0
						}
					} else {
						1
					}
				}
			} else {
				if inputs[7] <= 6 {
					if inputs[6] <= 95 {
						0
					} else {
						1
					}
				} else {
					1
				}
			}
		} else {
			0
		}
	} else {
		if inputs[7] <= 27 {
			if inputs[4] <= 52 {
				if inputs[11] <= 68 {
					1
				} else {
					if inputs[11] <= 187 {
						if inputs[10] <= 128 {
							0
						} else {
							if inputs[11] <= 159 {
								1
							} else {
								0
							}
						}
					} else {
						1
					}
				}
			} else {
				if inputs[0] <= 181 {
					if inputs[7] <= 4 {
						if inputs[0] <= 104 {
							1
						} else {
							0
						}
					} else {
						if inputs[6] <= 20 {
							if inputs[6] <= 15 {
								0
							} else {
								1
							}
						} else {
							if inputs[7] <= 21 {
								0
							} else {
								if inputs[4] <= 158 {
									0
								} else {
									if inputs[2] <= 5 {
										1
									} else {
										0
									}
								}
							}
						}
					}
				} else {
					if inputs[11] <= 101 {
						1
					} else {
						if inputs[3] <= 128 {
							0
						} else {
							1
						}
					}
				}
			}
		} else {
			if inputs[2] <= 2 {
				0
			} else {
				if inputs[11] <= 187 {
					if inputs[11] <= 96 {
						if inputs[4] <= 110 {
							if inputs[11] <= 88 {
								1
							} else {
								0
							}
						} else {
							0
						}
					} else {
						if inputs[8] <= 193 {
							1
						} else {
							if inputs[7] <= 36 {
								0
							} else {
								1
							}
						}
					}
				} else {
					if inputs[4] <= 33 {
						1
					} else {
						0
					}
				}
			}
		}
	}
}
