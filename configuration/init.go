package configuration

func Init() {
	setUpViper()
	registerDatabase()
	face_detection()
}
