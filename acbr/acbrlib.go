package acbr

/*
#cgo windows LDFLAGS: -L./lib -lACBrCEP
#cgo linux LDFLAGS: -L./lib -lacbrcep64

#include <stdlib.h>
#include <string.h>

// Declaração das funções da biblioteca
int CEP_Inicializar(const char* eArqConfig, const char* eChaveCrypt);
int CEP_Finalizar();
int CEP_BuscarPorCEP(const char* eCEP, char* buffer, int bufferLen);
*/
import "C"
import (
	"errors"
	"unsafe"
)

func Inicializar() error {
	res := C.CEP_Inicializar(nil, nil)
	if res != 0 {
		return errors.New("erro ao inicializar a biblioteca ACBr")
	}
	return nil
}

func Finalizar() {
	C.CEP_Finalizar()
}

func BuscarPorCEP(cep string) (string, error) {
	buffer := make([]byte, 1024) // Buffer para armazenar a resposta
	cCEP := C.CString(cep)
	defer C.free(unsafe.Pointer(cCEP))

	res := C.CEP_BuscarPorCEP(cCEP, (*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)))
	if res != 0 {
		return "", errors.New("erro ao buscar informações do CEP")
	}

	// Converter o buffer para string Go
	result := C.GoString((*C.char)(unsafe.Pointer(&buffer[0])))
	return result, nil
}
