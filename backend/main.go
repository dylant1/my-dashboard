package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

type Assignments []struct {
	ID                              int64     `json:"id"`
	Description                     string    `json:"description"`
	DueAt                           time.Time `json:"due_at"`
	UnlockAt                        any       `json:"unlock_at"`
	LockAt                          any       `json:"lock_at"`
	PointsPossible                  float64   `json:"points_possible"`
	GradingType                     string    `json:"grading_type"`
	AssignmentGroupID               int64     `json:"assignment_group_id"`
	GradingStandardID               any       `json:"grading_standard_id"`
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
	PeerReviews                     bool      `json:"peer_reviews"`
	AutomaticPeerReviews            bool      `json:"automatic_peer_reviews"`
	Position                        int       `json:"position"`
	GradeGroupStudentsIndividually  bool      `json:"grade_group_students_individually"`
	AnonymousPeerReviews            bool      `json:"anonymous_peer_reviews"`
	GroupCategoryID                 any       `json:"group_category_id"`
	PostToSis                       bool      `json:"post_to_sis"`
	ModeratedGrading                bool      `json:"moderated_grading"`
	OmitFromFinalGrade              bool      `json:"omit_from_final_grade"`
	IntraGroupPeerReviews           bool      `json:"intra_group_peer_reviews"`
	AnonymousInstructorAnnotations  bool      `json:"anonymous_instructor_annotations"`
	AnonymousGrading                bool      `json:"anonymous_grading"`
	GradersAnonymousToGraders       bool      `json:"graders_anonymous_to_graders"`
	GraderCount                     int       `json:"grader_count"`
	GraderCommentsVisibleToGraders  bool      `json:"grader_comments_visible_to_graders"`
	FinalGraderID                   any       `json:"final_grader_id"`
	GraderNamesVisibleToFinalGrader bool      `json:"grader_names_visible_to_final_grader"`
	AllowedAttempts                 int       `json:"allowed_attempts"`
	AnnotatableAttachmentID         any       `json:"annotatable_attachment_id"`
	HideInGradebook                 bool      `json:"hide_in_gradebook"`
	SecureParams                    string    `json:"secure_params"`
	LtiContextID                    string    `json:"lti_context_id"`
	CourseID                        int64     `json:"course_id"`
	Name                            string    `json:"name"`
	SubmissionTypes                 []string  `json:"submission_types"`
	HasSubmittedSubmissions         bool      `json:"has_submitted_submissions"`
	DueDateRequired                 bool      `json:"due_date_required"`
	MaxNameLength                   int       `json:"max_name_length"`
	InClosedGradingPeriod           bool      `json:"in_closed_grading_period"`
	GradedSubmissionsExist          bool      `json:"graded_submissions_exist"`
	IsQuizAssignment                bool      `json:"is_quiz_assignment"`
	CanDuplicate                    bool      `json:"can_duplicate"`
	OriginalCourseID                any       `json:"original_course_id"`
	OriginalAssignmentID            any       `json:"original_assignment_id"`
	OriginalLtiResourceLinkID       any       `json:"original_lti_resource_link_id"`
	OriginalAssignmentName          any       `json:"original_assignment_name"`
	OriginalQuizID                  any       `json:"original_quiz_id"`
	WorkflowState                   string    `json:"workflow_state"`
	ImportantDates                  bool      `json:"important_dates"`
	Muted                           bool      `json:"muted"`
	HTMLURL                         string    `json:"html_url"`
	Published                       bool      `json:"published"`
	OnlyVisibleToOverrides          bool      `json:"only_visible_to_overrides"`
	Bucket                          string    `json:"bucket"`
	LockedForUser                   bool      `json:"locked_for_user"`
	SubmissionsDownloadURL          string    `json:"submissions_download_url"`
	PostManually                    bool      `json:"post_manually"`
	AnonymizeStudents               bool      `json:"anonymize_students"`
	RequireLockdownBrowser          bool      `json:"require_lockdown_browser"`
	RestrictQuantitativeData        bool      `json:"restrict_quantitative_data"`
}
type Assignment struct {
	ID                              int64     `json:"id"`
	Description                     string    `json:"description"`
	DueAt                           time.Time `json:"due_at"`
	UnlockAt                        any       `json:"unlock_at"`
	LockAt                          any       `json:"lock_at"`
	PointsPossible                  float64   `json:"points_possible"`
	GradingType                     string    `json:"grading_type"`
	AssignmentGroupID               int64     `json:"assignment_group_id"`
	GradingStandardID               any       `json:"grading_standard_id"`
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
	PeerReviews                     bool      `json:"peer_reviews"`
	AutomaticPeerReviews            bool      `json:"automatic_peer_reviews"`
	Position                        int       `json:"position"`
	GradeGroupStudentsIndividually  bool      `json:"grade_group_students_individually"`
	AnonymousPeerReviews            bool      `json:"anonymous_peer_reviews"`
	GroupCategoryID                 any       `json:"group_category_id"`
	PostToSis                       bool      `json:"post_to_sis"`
	ModeratedGrading                bool      `json:"moderated_grading"`
	OmitFromFinalGrade              bool      `json:"omit_from_final_grade"`
	IntraGroupPeerReviews           bool      `json:"intra_group_peer_reviews"`
	AnonymousInstructorAnnotations  bool      `json:"anonymous_instructor_annotations"`
	AnonymousGrading                bool      `json:"anonymous_grading"`
	GradersAnonymousToGraders       bool      `json:"graders_anonymous_to_graders"`
	GraderCount                     int       `json:"grader_count"`
	GraderCommentsVisibleToGraders  bool      `json:"grader_comments_visible_to_graders"`
	FinalGraderID                   any       `json:"final_grader_id"`
	GraderNamesVisibleToFinalGrader bool      `json:"grader_names_visible_to_final_grader"`
	AllowedAttempts                 int       `json:"allowed_attempts"`
	AnnotatableAttachmentID         any       `json:"annotatable_attachment_id"`
	HideInGradebook                 bool      `json:"hide_in_gradebook"`
	SecureParams                    string    `json:"secure_params"`
	LtiContextID                    string    `json:"lti_context_id"`
	CourseID                        int64     `json:"course_id"`
	Name                            string    `json:"name"`
	SubmissionTypes                 []string  `json:"submission_types"`
	HasSubmittedSubmissions         bool      `json:"has_submitted_submissions"`
	DueDateRequired                 bool      `json:"due_date_required"`
	MaxNameLength                   int       `json:"max_name_length"`
	InClosedGradingPeriod           bool      `json:"in_closed_grading_period"`
	GradedSubmissionsExist          bool      `json:"graded_submissions_exist"`
	IsQuizAssignment                bool      `json:"is_quiz_assignment"`
	CanDuplicate                    bool      `json:"can_duplicate"`
	OriginalCourseID                any       `json:"original_course_id"`
	OriginalAssignmentID            any       `json:"original_assignment_id"`
	OriginalLtiResourceLinkID       any       `json:"original_lti_resource_link_id"`
	OriginalAssignmentName          any       `json:"original_assignment_name"`
	OriginalQuizID                  any       `json:"original_quiz_id"`
	WorkflowState                   string    `json:"workflow_state"`
	ImportantDates                  bool      `json:"important_dates"`
	Muted                           bool      `json:"muted"`
	HTMLURL                         string    `json:"html_url"`
	Published                       bool      `json:"published"`
	OnlyVisibleToOverrides          bool      `json:"only_visible_to_overrides"`
	Bucket                          string    `json:"bucket"`
	LockedForUser                   bool      `json:"locked_for_user"`
	SubmissionsDownloadURL          string    `json:"submissions_download_url"`
	PostManually                    bool      `json:"post_manually"`
	AnonymizeStudents               bool      `json:"anonymize_students"`
	RequireLockdownBrowser          bool      `json:"require_lockdown_browser"`
	RestrictQuantitativeData        bool      `json:"restrict_quantitative_data"`
}
type UserInfo []struct {
	ID                          int64     `json:"id"`
	AccessRestrictedByDate      bool      `json:"access_restricted_by_date,omitempty"`
	Name                        string    `json:"name,omitempty"`
	AccountID                   int64     `json:"account_id,omitempty"`
	UUID                        string    `json:"uuid,omitempty"`
	StartAt                     any       `json:"start_at,omitempty"`
	GradingStandardID           any       `json:"grading_standard_id,omitempty"`
	IsPublic                    any       `json:"is_public,omitempty"`
	CreatedAt                   time.Time `json:"created_at,omitempty"`
	CourseCode                  string    `json:"course_code,omitempty"`
	DefaultView                 string    `json:"default_view,omitempty"`
	RootAccountID               int64     `json:"root_account_id,omitempty"`
	EnrollmentTermID            int64     `json:"enrollment_term_id,omitempty"`
	License                     string    `json:"license,omitempty"`
	GradePassbackSetting        any       `json:"grade_passback_setting,omitempty"`
	EndAt                       any       `json:"end_at,omitempty"`
	PublicSyllabus              bool      `json:"public_syllabus,omitempty"`
	PublicSyllabusToAuth        bool      `json:"public_syllabus_to_auth,omitempty"`
	StorageQuotaMb              int       `json:"storage_quota_mb,omitempty"`
	IsPublicToAuthUsers         bool      `json:"is_public_to_auth_users,omitempty"`
	HomeroomCourse              bool      `json:"homeroom_course,omitempty"`
	CourseColor                 any       `json:"course_color,omitempty"`
	FriendlyName                any       `json:"friendly_name,omitempty"`
	ApplyAssignmentGroupWeights bool      `json:"apply_assignment_group_weights,omitempty"`
	Calendar                    struct {
		Ics string `json:"ics"`
	} `json:"calendar,omitempty"`
	TimeZone    string `json:"time_zone,omitempty"`
	Blueprint   bool   `json:"blueprint,omitempty"`
	Template    bool   `json:"template,omitempty"`
	Enrollments []struct {
		Type                           string `json:"type"`
		Role                           string `json:"role"`
		RoleID                         int64  `json:"role_id"`
		UserID                         int64  `json:"user_id"`
		EnrollmentState                string `json:"enrollment_state"`
		LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
	} `json:"enrollments,omitempty"`
	HideFinalGrades                  bool   `json:"hide_final_grades,omitempty"`
	WorkflowState                    string `json:"workflow_state,omitempty"`
	CourseFormat                     string `json:"course_format,omitempty"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates,omitempty"`
}

// Retrieve upcoming assignments due
func getAssignments() []Assignment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ACCESS_TOKEN := os.Getenv("ACCESS_TOKEN")
	url := fmt.Sprintf("https://canvas.instructure.com/api/v1/courses?access_token=%s", ACCESS_TOKEN)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	user := UserInfo{}
	json.Unmarshal(body, &user)

	course_ids := []int64{}

	for _, course := range user {
		if len(course.Name) > 0 {
			course_ids = append(course_ids, course.ID)
		}
	}
	assignments := []Assignment{}

	for i := 0; i < len(course_ids); i++ {
		url = fmt.Sprintf("https://canvas.instructure.com/api/v1/courses/%d/assignments?bucket=upcoming&access_token=%s", course_ids[i], ACCESS_TOKEN)
		resp, err = http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if string(body) != "[]" {
			assignment_list := Assignments{}
			json.Unmarshal(body, &assignment_list)

			for _, assignment := range assignment_list {
				if len(assignment.Name) > 0 {
					tmp := Assignment{}
					tmp = assignment
					assignments = append(assignments, tmp)
				}
			}
		}

	}

	return assignments
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func main() {

	c := cache.New(5*time.Minute, 10*time.Minute)

	mux := http.NewServeMux()

	mux.HandleFunc("/assignments", func(rw http.ResponseWriter, _ *http.Request) {
		enableCors(&rw)
		assignments, found := c.Get("assignments")
		if !found {
			assignments = getAssignments()
			c.Set("assignments", assignments, cache.DefaultExpiration)
		}
		rw.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(assignments)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(js)
		// json.NewEncoder(rw).Encode(assignments)

	})
	log.Println("Listening...")
	http.ListenAndServe("127.0.0.1:3000", mux)
}
