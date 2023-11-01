// Code generated from Python2Noir
use dep::std::println;
use dep::quantization_arithmetic;
fn obtainEuclideanDistance(inputs : [Field;13],inputs_scale : Field,inputs_zero_point : Field,point : [Field;13],point_scale : Field,point_zero_point : Field,) -> pub Field {
	let mut sum: Field = 0;
	for index in 0..5 {
		let x: Field = quantization_arithmetic::sub(point[index], point_zero_point, point_scale, inputs[index], inputs_zero_point, inputs_scale, point_zero_point, point_scale);
		let x_double: Field = quantization_arithmetic::mul(x, point_zero_point, point_scale, x, point_zero_point, point_scale, point_zero_point, point_scale);
		sum = quantization_arithmetic::sub(sum, point_zero_point, point_scale, x_double, point_zero_point, point_scale, point_zero_point, point_scale);
	}
	sum * 255
}
fn check_min(e0 : Field,e1 : Field,e2 : Field,e3 : Field,e4 : Field,) -> pub u3 {
	let mut output: u3 = 0;
	let mut temp: Field = e0;
	if e1 as u32 / 255 < temp as u32 / 255 {
		temp = e1;
		output = 1;
	}
	if e2 as u32 / 255 < temp as u32 / 255 {
		temp = e2;
		output = 2;
	}
	if e3 as u32 / 255 < temp as u32 / 255 {
		temp = e3;
		output = 3;
	}
	if e4 as u32 / 255 < temp as u32 / 255 {
		temp = e4;
		output = 4;
	}
	output
}
fn main(inputs : [Field;13],inputs_scale : Field,inputs_zero_point : Field,) -> pub u3 {
	let point_scale: Field = 256;
	let point_zero_point: Field = 0-1;
	let point0: [Field;13] = [156, 162, 123, 100, 76, 255, 162, 159, 34, 29, 63, 61, 19];
	let point1: [Field;13] = [134, 238, 207, 84, 64, 40, 0, 138, 131, 52, 92, 71, 248];
	let point2: [Field;13] = [122, 114, 157, 81, 65, 0, 1, 170, 29, 25, 45, 25, 2];
	let point3: [Field;13] = [152, 216, 221, 106, 76, 38, 251, 133, 139, 72, 111, 100, 226];
	let point4: [Field;13] = [125, 131, 170, 85, 74, 0, 253, 166, 53, 26, 63, 31, 2];
	let e0: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point0, point_scale, point_zero_point);
	let e1: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point1, point_scale, point_zero_point);
	let e2: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point2, point_scale, point_zero_point);
	let e3: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point3, point_scale, point_zero_point);
	let e4: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point4, point_scale, point_zero_point);
	check_min(e0,e1,e2,e3,e4)
}