/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: serializer.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-13 16:44:20 (CST)
 Last Update: Thursday 2018-09-13 16:54:17 (CST)
		  By:
 Description:
 *********************************************************************************/
package web

// SerializerType ..
type SerializerType interface {
	Serializer() map[string]interface{}
}

// Serializer ..
type Serializer struct {
	Instances []SerializerType
	Instance  SerializerType
}

// Data ..
func (s *Serializer) Data() []map[string]interface{} {
	ins := make([]map[string]interface{}, 0)
	if s.Instances != nil {
		for _, instance := range s.Instances {
			ins = append(ins, instance.Serializer())
		}
		return ins
	}
	ins = append(ins, s.Instance.Serializer())
	return ins
}
