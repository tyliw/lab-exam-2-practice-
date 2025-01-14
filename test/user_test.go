package unit

import (
	"fmt"
	"lab-exam/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)


func TestValidUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`user is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B6506469",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestInvalidUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`student_id is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"", 			// ผิดตรงนี้ StudentID ไม่มีต่า	
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",		
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run(`first_name is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"",					// ผิดตรงนี้ FirstName ไม่มีต่า
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})

	t.Run(`last_name is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"Ravipon",
			LastName: 	"",					// ผิดตรงนี้ LastName ไม่มีต่า
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("LastName is required"))
	})

	t.Run(`email is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",					
			Email: 		"",					// ผิดตรงนี้ Email ไม่มีต่า
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run(`phone is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",					
			Email: 		"test@gmail.com",
			Phone: 		"",					// ผิดตรงนี้ Phone ไม่มีต่า
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`profile is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",					
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",					
			Profile: 	"",					// Profile สามารถไม่มีต่าได้
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		
	})

	t.Run(`link_in is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",					
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",					
			Profile: 	"profile.pdf",					
			LinkIn: 	"",					// ผิดตรงนี้ LinkIn ไม่มีต่า
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("LinkIn is required"))
	})

	t.Run(`gender_id is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B1234567",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",					
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",					
			Profile: 	"profile.pdf",					
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	0,					// ผิดตรงนี้ GenderID ไม่มีต่า
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Gender is required"))
	})
}

func TestStudentIDFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`student alphabet incorrect format`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"A1234567",		// ผิดตรงนี้ StudentID มี format ไม่ถูกต้อง format ต้องขึ้นต้นด้วยตัวอักษร B, M หรือ D 1 ตัว และ ตามตัวเลข 7 ตัว (ผิดตัวขี้นต้น)
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID)))
	})

	t.Run(`student number incorrect format`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B123456",		// ผิดตรงนี้ StudentID มี format ไม่ถูกต้อง format ต้องขึ้นต้นด้วยตัวอักษร B, M หรือ D 1 ตัว และ ตามตัวเลข 7 ตัว (ผิดตัวเลขไม่ครบ 7 ตัว)
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID)))
	})
}

func TestEmailFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`email incorrect format`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B6506469",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test.com",			// ผิดตรงนี้ Email มี format ไม่ถูกต้อง format ต้องมีรูปแบบ email ที่ถูกต้อง
			Phone: 		"0979989859",
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}
		
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}

func TestPhoneFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`phone incorrect format`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B6506469",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"123456789",		// ผิดตรงนี้ Phone มี format ไม่ถูกต้อง format ต้องมีตัวเลขทั้งหมด 10 ตัว (ผิดตัวเลขไม่ครบ 10 ตัว)
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}
		
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone must be a 10-digit number"))
	})

	t.Run(`phone incorrect format`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B6506469",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"abcdefghij",		// ผิดตรงนี้ Phone มี format ไม่ถูกต้อง format ต้องมีตัวเลขทั้งหมด 10 ตัว (ผิดตัวตรงตัวอักษร)
			Profile: 	"profile.pdf",
			LinkIn: 	"https://www.linkedin.com/company/ilink/",
			GenderID: 	1,
		}
		
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone must be a 10-digit number"))
	})
}

func TestLinkInFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`link_in incorrect format`, func(t *testing.T) {
		user := entity.User{
			StudentID: 	"B6506469",
			FirstName: 	"Ravipon",
			LastName: 	"Mungdee",
			Email: 		"test@gmail.com",
			Phone: 		"0979989859",		
			Profile: 	"profile.pdf",
			LinkIn: 	"linkedin",		// ผิดตรงนี้ LinkIn มี format ไม่ถูกต้อง format ต้องเป็น link url ที่ถูกต้อง
			GenderID: 	1,
		}
		
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Url LinkIn is invalid"))
	})
}