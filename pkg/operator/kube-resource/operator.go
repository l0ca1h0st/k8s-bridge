package kube_resource

import "k8s.io/client-go/kubernetes"

type Operator interface {
	AddOperator()error
	UpdateOperator()error
	DeleteOperator()error
}

type kubeResourceOperator struct {
	clientSet kubernetes.Interface
}

func NewKubeResourceOperator(clientSet kubernetes.Interface)Operator{
	return &kubeResourceOperator{clientSet:clientSet}
}

func (op *kubeResourceOperator) AddOperator()error{
	return nil
}

func(op *kubeResourceOperator)UpdateOperator()error{
	return nil
}

func (op *kubeResourceOperator)DeleteOperator()error{
	return nil
}