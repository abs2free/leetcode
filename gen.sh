#!/bin/bash

echo "// package main this is a leetcode
package main

/*
*

*
 */
func $2(in string) string {
    return \"\"
}" > $1.go;


foo=$2;
upperName="$(tr '[:lower:]' '[:upper:]' <<< ${foo:0:1})${foo:1}"

echo "package main

import \"testing\"

var $2Cases = []struct {
	name   string
	input  string
	except string
}{
	{
	},
}

func Test$upperName(t *testing.T) {
	t.Parallel()
	for _, c := range $2Cases {
		t.Run(c.name, func(t *testing.T) {
			actual := $2(c.input)
			if actual != c.except {
				t.Errorf(\"$2 $s  has fail: input:%v ,except:%v, actual:%v \\\n\", c.name, c.input, c.except, actual)
			}
		})
	}
}" > "$1_test.go"
