// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
use dep::std::println;
use dep::quantization_arithmetic;
fn trees0class0(inputs : [u32;22],) -> pub u32 {
	if inputs[21] <= 47 {
		if inputs[13] <= 15 {
			if inputs[0] <= 131 {
				168
			} else {
				89
			}
		} else {
			68
		}
	} else {
		if inputs[11] <= 24 {
			if inputs[0] <= 44 {
				if inputs[0] <= 41 {
					127
				} else {
					81
				}
			} else {
				if inputs[10] <= 27 {
					192
				} else {
					118
				}
			}
		} else {
			200
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees0class1(inputs : [u32;22],) -> pub u32 {
	if inputs[18] <= 77 {
		if inputs[1] <= 66 {
			if inputs[17] <= 157 {
				170
			} else {
				87
			}
		} else {
			76
		}
	} else {
		if inputs[9] <= 22 {
			if inputs[0] <= 44 {
				98
			} else {
				if inputs[3] <= 27 {
					177
				} else {
					130
				}
			}
		} else {
			184
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees1class0(inputs : [u32;22],) -> pub u32 {
	if inputs[18] <= 77 {
		if inputs[1] <= 66 {
			if inputs[17] <= 157 {
				if inputs[21] <= 39 {
					168
				} else {
					137
				}
			} else {
				94
			}
		} else {
			84
		}
	} else {
		if inputs[9] <= 22 {
			if inputs[0] <= 44 {
				103
			} else {
				if inputs[8] <= 21 {
					170
				} else {
					127
				}
			}
		} else {
			176
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees1class1(inputs : [u32;22],) -> pub u32 {
	if inputs[21] <= 47 {
		if inputs[8] <= 4 {
			146
		} else {
			if inputs[19] <= 99 {
				85
			} else {
				if inputs[4] <= 18 {
					146
				} else {
					99
				}
			}
		}
	} else {
		if inputs[11] <= 24 {
			if inputs[0] <= 44 {
				if inputs[5] <= 13 {
					97
				} else {
					126
				}
			} else {
				if inputs[9] <= 21 {
					170
				} else {
					119
				}
			}
		} else {
			172
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees2class0(inputs : [u32;22],) -> pub u32 {
	if inputs[21] <= 47 {
		if inputs[16] <= 48 {
			141
		} else {
			if inputs[15] <= 181 {
				89
			} else {
				123
			}
		}
	} else {
		if inputs[11] <= 24 {
			if inputs[0] <= 44 {
				if inputs[8] <= 17 {
					100
				} else {
					127
				}
			} else {
				if inputs[10] <= 27 {
					165
				} else {
					115
				}
			}
		} else {
			168
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees2class1(inputs : [u32;22],) -> pub u32 {
	if inputs[21] <= 32 {
		97
	} else {
		if inputs[5] <= 15 {
			if inputs[0] <= 67 {
				if inputs[18] <= 91 {
					92
				} else {
					121
				}
			} else {
				160
			}
		} else {
			if inputs[17] <= 96 {
				if inputs[17] <= 82 {
					145
				} else {
					117
				}
			} else {
				if inputs[0] <= 159 {
					168
				} else {
					138
				}
			}
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees3class0(inputs : [u32;22],) -> pub u32 {
	if inputs[21] <= 56 {
		if inputs[1] <= 66 {
			if inputs[17] <= 159 {
				152
			} else {
				109
			}
		} else {
			94
		}
	} else {
		if inputs[16] <= 94 {
			123
		} else {
			if inputs[10] <= 23 {
				135
			} else {
				166
			}
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees3class1(inputs : [u32;22],) -> pub u32 {
	if inputs[19] <= 116 {
		if inputs[1] <= 56 {
			if inputs[0] <= 44 {
				113
			} else {
				158
			}
		} else {
			96
		}
	} else {
		if inputs[14] <= 4 {
			125
		} else {
			162
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees4class0(inputs : [u32;22],) -> pub u32 {
	if inputs[5] <= 15 {
		if inputs[0] <= 67 {
			100
		} else {
			if inputs[0] <= 155 {
				156
			} else {
				108
			}
		}
	} else {
		if inputs[17] <= 96 {
			126
		} else {
			if inputs[5] <= 49 {
				163
			} else {
				135
			}
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees4class1(inputs : [u32;22],) -> pub u32 {
	if inputs[19] <= 99 {
		if inputs[1] <= 55 {
			133
		} else {
			98
		}
	} else {
		if inputs[20] <= 91 {
			if inputs[18] <= 94 {
				113
			} else {
				140
			}
		} else {
			158
		}
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn count_prob<N>(class_prob : [u32;N],) -> pub u32 {
	let mut cls: Field = class_prob[0] as Field;
	for i in 1..N {
		cls = quantization_arithmetic::add(cls,127,128,class_prob[i] as Field,127,128,127,128);
	}
	cls = quantization_arithmetic::mul(cls,127,128,1,0,1,125,13);
	((cls*128) as u32)/128
}
fn main(inputs : [u32;22],) -> pub u32 {
	let trees0_class0: u32 = trees0class0(inputs);
	let trees0_class1: u32 = trees0class1(inputs);
	let trees1_class0: u32 = trees1class0(inputs);
	let trees1_class1: u32 = trees1class1(inputs);
	let trees2_class0: u32 = trees2class0(inputs);
	let trees2_class1: u32 = trees2class1(inputs);
	let trees3_class0: u32 = trees3class0(inputs);
	let trees3_class1: u32 = trees3class1(inputs);
	let trees4_class0: u32 = trees4class0(inputs);
	let trees4_class1: u32 = trees4class1(inputs);
	let c0: u32 = count_prob([trees0_class0, trees1_class0, trees2_class0, trees3_class0, trees4_class0]);
	let c1: u32 = count_prob([trees0_class1, trees1_class1, trees2_class1, trees3_class1, trees4_class1]);
	let mut max_ele_index: u32 = c0;
	let mut res: u32 = 0;
	if c1 > max_ele_index {
		max_ele_index = c1;
		res = 1;
	}
	println(res);
	res
}