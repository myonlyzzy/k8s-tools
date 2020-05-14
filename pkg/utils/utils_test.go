package utils

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
	"testing"
)

func TestRemovePod(t *testing.T) {
	type args struct {
		p []v1.Pod
		i int
	}
	pods := []v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod0",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod1",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod2",
			},
		},
	}
	want := []v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod0",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "pod2",
			},
		},
	}
	tests := []struct {
		name string
		args args
		want []v1.Pod
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				p: pods,
				i: 1,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemovePod(tt.args.p, tt.args.i)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovePod() = %v, want %v", got, tt.want)
			}
		})
	}
}
