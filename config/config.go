package config

type configurationStruct struct {
	Teste1 string
	teste2 string
}

var ConfigurationInstance configurationStruct
var configurationInstance configurationStruct

func Init(t1, t2 string) error {
	configurationInstance = configurationStruct{
		Teste1: t1,
		teste2: t2,
	}

	ConfigurationInstance = configurationStruct{
		Teste1: t1,
		teste2: t2,
	}

	return nil
}

func privateTest(t string) string {
	return t
}

func PublicTest(t string) string {
	return t
}
