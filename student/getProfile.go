package student

import (
	"database/sql"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getProfileResponse struct{
	FirstName string
	LastName string
	Gender string
	DOB string
	Email string
	GPA sql.NullFloat64
	SCU sql.NullInt32
	batch int
	Major string
}

func GetStudentProfile(c *fiber.Ctx) error {
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "s" {
		return c.SendStatus(403)
	}
	sqlResult, err := db.Db.Query(
		"select first_name,last_name,gender,dob,email, (select sum(scu) from Courses c,Enrolled_Courses e, Class cl where e.student_id = (select s.student_id from Students s where s.user_id = ?) and e.class_id = cl.class_id and cl.course_id = c.course_id) as scu ,batch,m.major_name from Students s,Majors m where user_id = ? and m.major_id = s.major_id",id,id)
	if err != nil {
		log.Println("-1")
		log.Println(err)
		return c.SendStatus(500)
	}
	if sqlResult.Err() != nil {
		log.Println("0")
		log.Println(err)
		return c.SendStatus(500)
	}
	var r getProfileResponse
	f := false
	for sqlResult.Next() {
		f = true
		err = sqlResult.Scan(&r.FirstName,&r.LastName,&r.Gender,&r.DOB,&r.Email,&r.SCU,&r.batch,&r.Major)
		if err != nil {
			log.Println(err)
			log.Println("6")
			return c.SendStatus(500)
		}
	}
	if !f {
		return c.SendStatus(404)
	}
	res2, err := db.Db.Query(
		"select e.mid_score,e.final_score from Enrolled_Courses e where e.student_id = (select student_id from Students s where s.user_id = ?)",id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	scores := []float64{}
	count := 0
	v := false
	for res2.Next() {
		v = true
		var mid int
		var final int
		err = res2.Scan(&mid,&final)
		if err != nil {
			return c.SendStatus(500)
		}
		if mid < 0 || final < 0{
			v = false
			break
		}
		count++
		scores = append(scores, ((float64(mid)*0.3) + (float64(final)*0.7))*0.04)
	}
	if res2.Err() != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	if !v {
		r.GPA = sql.NullFloat64{
			Float64: 0.01,
			Valid:   false,
		}
		return c.JSON(r)
	}
	gpa := 0.0
	for _,score:=range scores{
		gpa += score
	}
	r.GPA = sql.NullFloat64{
		Float64: gpa/float64(count),
		Valid:   true,
	}
	return c.JSON(r)
}