export { Response, ErrorResponse, User, Token, Notification }

declare global {
  interface Response<T> {
    data: T,
  }

  interface ErrorResponse {
    message: string,
    status: int,
  }

  interface User {
    id: string,
    email: string,
    username: string,
    password: string,
    created_at: Date,
    updated_at: Date | undefined,
  }

  interface Token {
    id: string,
    user_id: string,
    value: string,
    expires_at: Date,
    created_at: Date,
    updated_at: Date | undefined,
  }
}
