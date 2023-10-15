// Code generated from Python2Noir
use dep::std::println;
mod post_quantization_operation;
fn obtainEuclideanDistance(inputs : [Field;22],inputs_scale : Field,inputs_zero_point : Field,point : [Field;22],point_scale : Field,point_zero_point : Field,) -> pub Field {
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
fn main(inputs : [Field;22],inputs_scale : Field,inputs_zero_point : Field,) -> pub u3 {
	let point_scale: Field = 364;
	let point_zero_point: Field = 0-13;
	let point0: [Field;22] = [87, 40, 61, 78, 79, 67, 75, 67, 113, 99, 130, 101, 75, 130, 40, 131, 255, 222, 201, 210, 172, 171];
	let point1: [Field;22] = [151, 69, 113, 16, 16, 12, 13, 12, 18, 14, 25, 14, 9, 25, 0, 219, 154, 179, 98, 141, 124, 72];
	let e0: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point0, point_scale, point_zero_point);
	let e1: Field = obtainEuclideanDistance(inputs, inputs_scale, inputs_zero_point, point1, point_scale, point_zero_point);
	check_min(e0,e1)
}