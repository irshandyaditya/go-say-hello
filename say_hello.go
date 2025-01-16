package gosayhello

// membuat module dan up ke github
func SayHello(name string) string {
	return "Heyyy, " + name
}

// major changes lebih baik mengubah nama module agar ketika go get dilakukan,
// pemakainya tidak mengalami error pada semua penggunaan yang mengalami perubahan
