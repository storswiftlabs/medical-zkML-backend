// Code generated from Python2Noir
use dep::std::println;
mod post_quantization_operation;
fn obtainEuclideanDistance(inputs : [Field;12],inputs_scale : Field,inputs_zero_point : Field,point : [Field;12],point_scale : Field,point_zero_point : Field,) -> pub Field {
	let mut sum: Field = 0;
	for index in 0..5 {
		let x: Field = post_quantization_operation::sub(point[index], point_zero_point, point_scale, inputs[index], inputs_zero_point, inputs_scale, point_zero_point, point_scale);
		let x_double: Field = post_quantization_operation::mul(x, point_zero_point, point_scale, x, point_zero_point, point_scale, point_zero_point, point_scale);
		sum = post_quantization_operation::sub(sum, point_zero_point, point_scale, x_double, point_zero_point, point_scale, point_zero_point, point_scale);
	}
	sum * 255
}
fn check_min(e0 : Field,e1 : Field,) -> pub u3 {
	let mut output: u3 = 0;
	let mut temp: Field = e0;
	if e1 as u32 / 255 < temp as u32 / 255 {
		temp = e1;
		output = 1;
	}
	output
}
fn main(inputs : [Field;12],inputs_scale : Field,inputs_zero_point : Field,) -> pub u3 {
	let point_scale: Field = 263;
	let point_zero_point: Field = 0-5;
	let point0: [Field;12] = [101, 99, 15, 43, 85, 84, 68, 21, 173, 255, 156, 110];
	let point1: [Field;12] = [88, 118, 13, 171, 97, 91, 74, 22, 172, 71, 0, 116];
	let e0: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point0, point_scale, point_zero_point);
	let e1: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point1, point_scale, point_zero_point);
	check_min(e0,e1)
}