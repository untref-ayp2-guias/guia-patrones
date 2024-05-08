package ejercicios

type FileMp3 struct {
	name string
}

func (f *FileMp3) PlayMp3() string {
	return "Reproduciendo archivo MP3. Nombre: " + f.name
}

func NewFileMp3(name string) *FileMp3 {
	return &FileMp3{name: name}
}

type ArchivoDeMusica interface {
	Reproducir() string
}