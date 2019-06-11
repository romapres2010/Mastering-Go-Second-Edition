package main

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func Add(sum_arg1, sum_arg2 int8) (interface{}, error) {
	sum_scope := op.NewScope()
	input1 := op.Placeholder(sum_scope.SubScope("Summand1"), tf.Int8)
	input2 := op.Placeholder(sum_scope.SubScope("Summand2"), tf.Int8)
	sum_result_node := op.Add(sum_scope, input1, input2)

	graph, err := sum_scope.Finalize()

	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	summand1, err := tf.NewTensor(sum_arg1)
	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	summand2, err := tf.NewTensor(sum_arg2)
	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	session, err := tf.NewSession(graph, nil)

	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}
	defer session.Close()

	sum, err := session.Run(
		map[tf.Output]*tf.Tensor{
			input1: summand1,
			input2: summand2,
		},
		[]tf.Output{sum_result_node}, nil)

	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	return sum[0].Value(), nil
}

func Multiply(sum_arg1, sum_arg2 int8) (interface{}, error) {
	sum_scope := op.NewScope()
	input1 := op.Placeholder(sum_scope.SubScope("Summand1"), tf.Int8)
	input2 := op.Placeholder(sum_scope.SubScope("Summand2"), tf.Int8)

	sum_result_node := op.Mul(sum_scope, input1, input2)
	graph, err := sum_scope.Finalize()

	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	summand1, err := tf.NewTensor(sum_arg1)
	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	summand2, err := tf.NewTensor(sum_arg2)
	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	session, err := tf.NewSession(graph, nil)

	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}
	defer session.Close()

	sum, err := session.Run(
		map[tf.Output]*tf.Tensor{
			input1: summand1,
			input2: summand2,
		},
		[]tf.Output{sum_result_node}, nil)

	if nil != err {
		fmt.Println(err.Error())
		return 0, err
	}

	return sum[0].Value(), nil
}

func main() {
	res, err := Add(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add:", res)
	}

	res, err = Multiply(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Multiply:", res)
	}
}
