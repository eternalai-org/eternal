package manager

// func Test_setupCondaEnv(t *testing.T) {
// 	type args struct {
// 		name    string
// 		envFile string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "Test setupCondaEnv",
// 			args: args{
// 				name:    "test1",
// 				envFile: "./env.yml",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := setupCondaEnv(tt.args.name, tt.args.envFile); (err != nil) != tt.wantErr {
// 				t.Errorf("setupCondaEnv() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestModelInstance_Setup(t *testing.T) {
// 	type fields struct {
// 		ModeID       string
// 		ModelAddress string
// 		ModelType    eaimodel.ModelType
// 		ModelPath    string
// 		ModelUrl     string
// 		Port         string
// 		inferLock    sync.Mutex
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		{
// 			name: "setup for IPFS (filecoin - lighthouse)",
// 			fields: fields{
// 				ModeID:       "1",
// 				ModelAddress: "0x123",
// 				ModelType:    "image",
// 				ModelPath:    "../tmp/1",
// 				ModelUrl:     "ipfs://QmX7j7FP2iwiQjsYPATRdYs8v9PGSVMpycxhsrxA1yHQAk",
// 				Port:         "9001",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &ModelInstance{
// 				ModelInfo: eaimodel.ModelInfoContract{
// 					ModelID:   big.NewInt(1),
// 					ModelAddr: tt.fields.ModelAddress,
// 					Metadata: eaimodel.ModelMetadata{
// 						ModelType:     tt.fields.ModelType,
// 						ModelURL:      tt.fields.ModelUrl,
// 						ModelFileHash: "",
// 						Version:       1,
// 						ModelName:     "test",
// 					},
// 				},
// 				ModelPath: tt.fields.ModelPath,
// 				Port:      tt.fields.Port,
// 			}
// 			if err := m.Setup(); (err != nil) != tt.wantErr {
// 				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
