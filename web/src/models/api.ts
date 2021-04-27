type ApiError = {
  code: number;
  error: string;
  message: string;
};

export type ApiResponse<T> = {
  response?: T;
  error?: ApiError;
};
