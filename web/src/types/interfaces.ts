export interface FormValue<T> {
  value: T,
  error?: string | null,
};

export interface User {
  id: string,
  email: string,
  first_name: string,
  last_name: string
};

export interface Feature {
  id: string,
  user_id: string,
  sprint_id: string,
  state: number,
  title: string,
  description: string,
  created_at: string,
}

export interface Requirement {
  id: string,
  user_id: string,
  feature_id: string,
  sprint_id: string,
  state: number,
  title: string,
  estimate: number,
  description: string,
  created_at: string,
}

export interface Response<T> {
  data: T
}
