package render

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	districtNumber = []string{
		"010", "020", "021", "022", "023", "024", "025", "027", "028", "029", "0310", "0311", "0312", "0313", "0314",
		"0315", "0316", "0317", "0318", "0319", "0335", "0350", "0351", "0352", "0353", "0354", "0355", "0356", "0357",
		"0358", "0359", "0370", "0371", "0372", "0373", "0374", "0375", "0376", "0377", "0378", "0379", "0391", "0392",
		"0393", "0394", "0395", "0396", "0398", "0410", "0411", "0412", "0413", "0414", "0415", "0416", "0417", "0418",
		"0419", "0421", "0427", "0429", "0431", "0432", "0433", "0434", "0435", "0436", "0437", "0438", "0439", "0440",
		"0450", "0451", "0452", "0453", "0454", "0455", "0456", "0457", "0458", "0459", "0470", "0471", "0472", "0473",
		"0474", "0475", "0476", "0477", "0478", "0479", "0482", "0483", "0510", "0511", "0512", "0513", "0514", "0515",
		"0516", "0517", "0518", "0519", "0523", "0530", "0531", "0532", "0533", "0534", "0535", "0536", "0537", "0538",
		"0539", "0550", "0551", "0552", "0553", "0554", "0555", "0556", "0557", "0558", "0559", "0561", "0562", "0563",
		"0564", "0565", "0566", "0570", "0571", "0572", "0573", "0574", "0575", "0576", "0577", "0578", "0579", "0580",
		"0591", "0592", "0593", "0594", "0595", "0596", "0597", "0598", "0599", "0660", "0661", "0662", "0663", "0691",
		"0692", "0701", "0710", "0711", "0712", "0713", "0714", "0715", "0716", "0717", "0718", "0719", "0722", "0724",
		"0728", "0730", "0731", "0732", "0733", "0734", "0735", "0736", "0737", "0738", "0739", "0743", "0744", "0745",
		"0746", "0751", "0752", "0753", "0754", "0755", "0756", "0757", "0758", "0759", "0760", "0762", "0763", "0765",
		"0766", "0768", "0769", "0770", "0771", "0772", "0773", "0774", "0775", "0776", "0777", "0778", "0779", "0790",
		"0791", "0792", "0793", "0794", "0795", "0796", "0797", "0798", "0799", "0810", "0811", "0812", "0813", "0814",
		"0816", "0817", "0818", "0819", "0825", "0826", "0827", "0830", "0831", "0832", "0833", "0834", "0835", "0836",
		"0837", "0838", "0839", "0840", "0851", "0852", "0853", "0854", "0855", "0856", "0857", "0858", "0859", "0870",
		"0871", "0872", "0873", "0874", "0875", "0876", "0877", "0878", "0879", "0881", "0883", "0886", "0887", "0888",
		"0890", "0891", "0892", "0893", "0898", "0899", "0910", "0911", "0912", "0913", "0914", "0915", "0916", "0917",
		"0919", "0930", "0931", "0932", "0933", "0934", "0935", "0936", "0937", "0938", "0941", "0943", "0951", "0952",
		"0953", "0954", "0971", "0972", "0973", "0974", "0975", "0976", "0977",
	}
)

func Phone(_ *MockContext) interface{} {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	phone := make([]string, 0)
	if time.Now().Unix()%2 == 0 {
		phone = append(phone, "5")
	} else {
		phone = append(phone, "6")
	}
	list := r.Perm(6)
	for _, item := range list {
		phone = append(phone, fmt.Sprintf("%v", item))
	}
	districtLength := len(districtNumber)
	dm := districtNumber[r.Int31n(int32(districtLength))]
	return "(" + dm + ")" + strings.Join(phone, "")
}

func Mobile(ctx *MockContext) interface{} {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	mobile := make([]string, 0)
	mobile = append(mobile, "1")
	list := r.Perm(10)
	for index, item := range list {
		if index == 0 && (item < 3 || item == 9 || item == 6) {
			ctx.r = [2]int64{1, 9}
			v := RangeInt(ctx)
			mobile = append(mobile, fmt.Sprintf("%v", v))
			continue
		}
		mobile = append(mobile, fmt.Sprintf("%v", item))
	}
	return "(+86)" + strings.Join(mobile, "")
}
