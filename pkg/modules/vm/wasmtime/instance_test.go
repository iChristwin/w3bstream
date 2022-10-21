package wasmtime_test

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	"github.com/iotexproject/w3bstream/pkg/modules/vm/common"
	"github.com/iotexproject/w3bstream/pkg/modules/vm/wasmtime"
	"github.com/iotexproject/w3bstream/pkg/types"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
)

// //go:embed ../../../../examples/log/log.wasm
var wasmLogCode []byte

// //go:embed ../../../examples/gjson/gjson.wasm
var wasmGJsonCode []byte

// //go:embed ../../../examples/easyjson/easyjson.wasm
var wasmEasyJsonCode []byte

// //go:embed ../../../../examples/word_count/word_count.wasm
var wasmWordCountCode []byte

// //go:embed ../../../../examples/word_count_v2/word_count_v2.wasm
var wasmWordCountV2Code []byte

func init() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	root := filepath.Join(wd, "../../../../examples")
	fmt.Println(root)

	var err error
	wasmLogCode, err = os.ReadFile(filepath.Join(root, "log/log.wasm"))
	if err != nil {
		panic(err)
	}

	wasmGJsonCode, err = os.ReadFile(filepath.Join(root, "gjson/gjson.wasm"))
	if err != nil {
		panic(err)
	}
	wasmEasyJsonCode, err = os.ReadFile(filepath.Join(root, "easyjson/easyjson.wasm"))
	if err != nil {
		panic(err)
	}
	wasmWordCountCode, err = os.ReadFile(filepath.Join(root, "word_count/word_count.wasm"))
	if err != nil {
		panic(err)
	}
	wasmWordCountV2Code, err = os.ReadFile(filepath.Join(root, "word_count_v2/word_count_v2.wasm"))
	if err != nil {
		panic(err)
	}
}

func TestInstance_LogWASM(t *testing.T) {
	require := require.New(t)
	i, err := wasmtime.NewInstanceByCode(context.Background(), wasmLogCode)
	require.NoError(err)
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	require.NoError(err)
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte("IoTeX"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	_, code = i.HandleEvent("not_exported", []byte("IoTeX"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_UnexportedHandler))
}

func TestInstance_GJsonWASM(t *testing.T) {
	require := require.New(t)
	i, err := wasmtime.NewInstanceByCode(context.Background(), wasmGJsonCode)
	require.NoError(err)
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	require.NoError(err)
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte(`{
  "name": {"first": "Tom", "last": "Anderson", "age": 39},
  "friends": [
    {"first_name": "Dale", "last_name": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first_name": "Roger", "last_name": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first_name": "Jane", "last_name": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))
}

func TestInstance_EasyJsonWASM(t *testing.T) {
	require := require.New(t)
	i, err := wasmtime.NewInstanceByCode(context.Background(), wasmEasyJsonCode)
	require.NoError(err)
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	require.NoError(err)
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte(`{"id":11,"student_name":"Tom","student_school":
								{"school_name":"MIT","school_addr":"xz"},
								"birthday":"2017-08-04T20:58:07.9894603+08:00"}`))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))
}

func TestInstance_WordCount(t *testing.T) {
	require := require.New(t)
	i, err := wasmtime.NewInstanceByCode(context.Background(), wasmWordCountCode)
	require.NoError(err)
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	require.NoError(err)
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte("a b c d a"))
	require.Equal(wasm.ResultStatusCode_OK, code)

	require.Equal(int32(2), i.Get("a"))
	require.Equal(int32(1), i.Get("b"))
	require.Equal(int32(1), i.Get("c"))
	require.Equal(int32(1), i.Get("d"))

	_, code = i.HandleEvent("start", []byte("a b c d a"))
	require.Equal(wasm.ResultStatusCode_OK, code)

	require.Equal(int32(4), i.Get("a"))
	require.Equal(int32(2), i.Get("b"))
	require.Equal(int32(2), i.Get("c"))
	require.Equal(int32(2), i.Get("d"))
}

func TestInstance_WordCountV2(t *testing.T) {
	require := require.New(t)
	i, err := wasmtime.NewInstanceByCode(context.Background(), wasmWordCountV2Code)
	require.NoError(err)
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	require.NoError(err)
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("count", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(i.Get("a")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("b")).To(Equal(int32(1)))
	NewWithT(t).Expect(i.Get("c")).To(Equal(int32(1)))
	NewWithT(t).Expect(i.Get("d")).To(Equal(int32(1)))

	_, code = i.HandleEvent("count", []byte("a b c d a"))
	NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))

	NewWithT(t).Expect(i.Get("a")).To(Equal(int32(4)))
	NewWithT(t).Expect(i.Get("b")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("c")).To(Equal(int32(2)))
	NewWithT(t).Expect(i.Get("d")).To(Equal(int32(2)))

	_, unique := i.HandleEvent("unique", nil)
	NewWithT(t).Expect(unique).To(Equal(wasm.ResultStatusCode(4)))
}

func TestInstance_TokenDistribute(t *testing.T) {
	i, err := wasmtime.NewInstanceByCode(context.Background(), wasmTokenDistributeCode)
	NewWithT(t).Expect(err).To(BeNil())
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	NewWithT(t).Expect(err).To(BeNil())
	defer vm.StopInstance(id)

	for idx := int32(0); idx < 20; idx++ {
		_, code := i.HandleEvent("start", []byte("test"))
		NewWithT(t).Expect(code).To(Equal(wasm.ResultStatusCode_OK))
		NewWithT(t).Expect(i.Get("clicks")).To(Equal(idx + 1))
	}
}
func TestInstance_SentTx(t *testing.T) {
	require := require.New(t)
	wasmSentTxCode, err := os.ReadFile("/home/haaai/iotex/tmp/test4/wasm4/lib6.wasm")
	require.NoError(err)

	data1 := `
	{
		"message":{
		   "steps":221,
		   "timestamp":1666118866
		},
		"signature":"c8bca57ba6f0e7936400069ce1681b2bfe2f52fd1f0a1e36f4a4762a89acb3f79a24264aca60215eb7b44e6de9cb97755bc3ef0d4884605780ad8471ccd31f10",
		"publicKey":"04b687e298ad52eec4fe32b27af45247f3659062593bd373c6ebb429ec840c90a906c1283f6b23a26e05a050d67d2215677781baaaa9e43a38d543052cec0c6afb",
		"deviceId":"b687e298ad52eec4fe32b27af45247f365906259",
		"cryptography":"ECC",
		"curve":"secp256r1",
		"hash":"sha256"
	 }
	`
	data2 := `
	{
		"message":{
		   "steps":239,
		   "timestamp":1666118872
		},
		"signature":"4dbe7d4ad4c4299f8624a902db5ec0c5232dc4901ce1ddc5c6e551fcbfb236a8855d1ac011b8b1b1744e5f48bf487736e235dfcb034a329b9b86d42426b49600",
		"publicKey":"04b687e298ad52eec4fe32b27af45247f3659062593bd373c6ebb429ec840c90a906c1283f6b23a26e05a050d67d2215677781baaaa9e43a38d543052cec0c6afb",
		"deviceId":"b687e298ad52eec4fe32b27af45247f365906259",
		"cryptography":"ECC",
		"curve":"secp256r1",
		"hash":"sha256"
	 }
	`
	data3 := `
	{
		"message":{
		   "steps":256,
		   "timestamp":1666118878
		},
		"signature":"80dc7a9fb7d4e49e50320477f39add261a42bee8b46f15d947534100bbffd0b5c01e38b7cb1284635e46a76faa66bf2feff035267e9f72056b035f2eb7d3e53a",
		"publicKey":"04b687e298ad52eec4fe32b27af45247f3659062593bd373c6ebb429ec840c90a906c1283f6b23a26e05a050d67d2215677781baaaa9e43a38d543052cec0c6afb",
		"deviceId":"b687e298ad52eec4fe32b27af45247f365906259",
		"cryptography":"ECC",
		"curve":"secp256r1",
		"hash":"sha256"
	 }
	`

	queryData := `
	{
		"blockHash":"0x2b5e0f9ba3c9fa11f6e1e3b0c09c557602b787cac6c96f6c7ff82ab89f65a865",
		"transactionHash":"0x504e8720229488a07bdd77397afb3524744f586fc5840c03e1dea9f4b99146e8",
		"logIndex":"0x0",
		"blockNumber":"0x1017c97",
		"transactionIndex":"0x1",
		"address":"0x1ea33901e0f9e0249881a7f71022b5fcee04d3b7",
		"data":"0x0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000169dc1cfc4fd15ed5276b12d6c10ce65fbef0d11000000000000000000000000b687e298ad52eec4fe32b27af45247f36590625900000000000000000000000000000000000000000000000000000000634ec98400000000000000000000000000000000000000000000000000000000634ef55d",
		"topics":[
		   "0x766e6460a49ca518797200f8d2b455a80962f1e6acdcda61000fc3dc2004db88"
		]
	 }
	`

	ctx := types.WithETHClientConfig(context.Background(), &wasm.ETHClientConfig{
		PrivateKey:    "",
		ChainEndpoint: "https://babel-api.testnet.iotex.io",
	})

	i, err := wasmtime.NewInstanceByCode(ctx,
		wasmSentTxCode, common.DefaultInstanceOptionSetter)
	require.NoError(err)
	id := vm.AddInstance(i)

	err = vm.StartInstance(id)
	require.NoError(err)
	defer vm.StopInstance(id)

	_, code := i.HandleEvent("start", []byte(data1))
	require.Equal(wasm.ResultStatusCode_OK, code)
	_, code = i.HandleEvent("start", []byte(data2))
	require.Equal(wasm.ResultStatusCode_OK, code)
	_, code = i.HandleEvent("start", []byte(data3))
	require.Equal(wasm.ResultStatusCode_OK, code)

	_, code = i.HandleEvent("claim", []byte(queryData))
	require.Equal(wasm.ResultStatusCode_OK, code)

}
