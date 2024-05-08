package types

type CompilationArgs struct {
	Version            string
	PathToMainContract string
	MainContractName   string
	BasePath           string
	SolcBinaryPath     string
	ContractAddress    string
	Optimization       bool
	Network            string
}

type ApiResponse struct {
	Success bool
	Result  string
}
