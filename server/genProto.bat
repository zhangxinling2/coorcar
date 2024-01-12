﻿set PROTO_PATH=.\auth\api
set GO_OUT_PATH=.\auth\api\gen\v1
mkdir  %GO_OUT_PATH%
protoc -I=%PROTO_PATH% --go_out=paths=source_relative:%GO_OUT_PATH% auth.proto
protoc -I=%PROTO_PATH% --go-grpc_out=paths=source_relative:%GO_OUT_PATH% auth.proto
protoc -I=%PROTO_PATH% --grpc-gateway_out=paths=source_relative,grpc_api_configuration=%PROTO_PATH%\auth.yaml:%GO_OUT_PATH% auth.proto
set PBTS_BIN_DIR=..\wx\miniprogram\node_modules\.bin
set PBTS_OUT_DIR=..\wx\miniprogram\service\proto_gen
mkdir  %PBTS_OUT_DIR%
start %PBTS_BIN_DIR%\pbjs -t static -w es6 %PROTO_PATH%\auth.proto --no-create --no-encode --no-decode --no-verify --no-delimited -o %PBTS_OUT_DIR%\auth_pb_tmp.js
echo import * as ^$protobuf from "protobufjs"; > %PBTS_OUT_DIR%\auth_pb.js
type %PBTS_OUT_DIR%\auth_pb_tmp.js >> %PBTS_OUT_DIR%\auth_pb.js
del %PBTS_OUT_DIR%\auth_pb_tmp.js
start %PBTS_BIN_DIR%\pbts -o %PBTS_OUT_DIR%\auth_pb.d.ts %PBTS_OUT_DIR%\auth_pb.js