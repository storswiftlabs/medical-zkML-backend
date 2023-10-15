// Code generated from Python2Noir
use dep::std::println;
mod post_quantization_operation;
fn obtainEuclideanDistance(inputs : [Field;18],inputs_scale : Field,inputs_zero_point : Field,point : [Field;18],point_scale : Field,point_zero_point : Field,) -> pub Field {
	let mut sum: Field = 0;
	for index in 0..5 {
		let x: Field = post_quantization_operation::sub(point[index], point_zero_point, point_scale, inputs[index], inputs_zero_point, inputs_scale, point_zero_point, point_scale);
		let x_double: Field = post_quantization_operation::mul(x, point_zero_point, point_scale, x, point_zero_point, point_scale, point_zero_point, point_scale);
		sum = post_quantization_operation::sub(sum, point_zero_point, point_scale, x_double, point_zero_point, point_scale, point_zero_point, point_scale);
	}
	sum * 255
}
fn check_min(e0 : Field,e1 : Field,e2 : Field,e3 : Field,) -> pub u3 {
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
	output
}
fn main(inputs : [Field;18],inputs_scale : Field,inputs_zero_point : Field,) -> pub u3 {
	let point_scale: Field = 255;
	let point_zero_point: Field = 0;
	let point0: [Field;18] = [128, 0, 0, 0, 26, 34, 26, 93, 21, 68, 132, 136, 116, 138, 94, 136, 128, 19];
	let point1: [Field;18] = [164, 255, 119, 42, 172, 249, 36, 202, 9, 142, 208, 200, 170, 157, 187, 243, 249, 80];
	let point2: [Field;18] = [142, 255, 41, 0, 14, 83, 0, 152, 0, 103, 186, 145, 159, 135, 148, 69, 152, 18];
	let point3: [Field;18] = [152, 13, 0, 0, 13, 114, 7, 248, 3, 172, 174, 177, 159, 180, 232, 208, 255, 104];
	let e0: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point0, point_scale, point_zero_point);
	let e1: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point1, point_scale, point_zero_point);
	let e2: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point2, point_scale, point_zero_point);
	let e3: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point3, point_scale, point_zero_point);
	check_min(e0,e1,e2,e3)
}