export type DiscountType = 'AMOUNT' | 'PERCENTAGE';

export interface ApplyGiftCodeRequest {
  code: string;
  userId: string;
}

export interface ApplyGiftCodeResponse {
  status: string;
  message: string;
  appliedValue: number;
  newBalance?: number;
}

export interface ErrorResponse {
  error: {
    code: string;
    message: string;
  };
}
