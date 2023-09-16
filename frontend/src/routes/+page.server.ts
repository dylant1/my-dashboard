/** @type {import('./$types').PageLoad} */

export type Assignments = Assignment[]

export interface Assignment {
  id: number
  description: string
  due_at: string
  unlock_at?: string
  lock_at?: string
  points_possible: number
  grading_type: string
  assignment_group_id: number
  grading_standard_id: any
  created_at: string
  updated_at: string
  peer_reviews: boolean
  automatic_peer_reviews: boolean
  position: number
  grade_group_students_individually: boolean
  anonymous_peer_reviews: boolean
  group_category_id: any
  post_to_sis: boolean
  moderated_grading: boolean
  omit_from_final_grade: boolean
  intra_group_peer_reviews: boolean
  anonymous_instructor_annotations: boolean
  anonymous_grading: boolean
  graders_anonymous_to_graders: boolean
  grader_count: number
  grader_comments_visible_to_graders: boolean
  final_grader_id: any
  grader_names_visible_to_final_grader: boolean
  allowed_attempts: number
  annotatable_attachment_id: any
  hide_in_gradebook: boolean
  secure_params: string
  lti_context_id: string
  course_id: number
  name: string
  submission_types: string[]
  has_submitted_submissions: boolean
  due_date_required: boolean
  max_name_length: number
  in_closed_grading_period: boolean
  graded_submissions_exist: boolean
  is_quiz_assignment: boolean
  can_duplicate: boolean
  original_course_id: any
  original_assignment_id: any
  original_lti_resource_link_id: any
  original_assignment_name: any
  original_quiz_id: any
  workflow_state: string
  important_dates: boolean
  muted: boolean
  html_url: string
  published: boolean
  only_visible_to_overrides: boolean
  bucket: string
  locked_for_user: boolean
  submissions_download_url: string
  post_manually: boolean
  anonymize_students: boolean
  require_lockdown_browser: boolean
  restrict_quantitative_data: boolean
}



export async function load({ params }) {

    let assignments: Assignments;

    const res = await  fetch("http://127.0.0.1:3000/assignments");

    assignments = await res.json();

    if (assignments) {
        return {
            assignments: assignments
        };
    }
	
}



