// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
use dep::std::println;
use dep::quantization_arithmetic;
fn trees0class0(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		53
	} else {
		195
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees0class1(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		70
	} else {
		181
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees1class0(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		78
	} else {
		173
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees1class1(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		82
	} else {
		168
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees2class0(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		86
	} else {
		164
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees2class1(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		88
	} else {
		161
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees3class0(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		90
	} else {
		159
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees3class1(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		92
	} else {
		156
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees4class0(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		94
	} else {
		154
	}
	// inputs quantization scale reciprocal: 128, zero-point: 127

}
fn trees4class1(inputs : [u32;24],) -> pub u32 {
	if inputs[3] <= 32 {
		96
	} else {
		152
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
fn main(inputs : [u32;24],) -> pub u32 {
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