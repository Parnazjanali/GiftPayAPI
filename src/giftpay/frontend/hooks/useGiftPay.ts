import { useState } from 'react';
import { ApplyGiftCodeRequest, ApplyGiftCodeResponse, ErrorResponse } from '../types/giftpay';

export const useGiftPay = (apiUrl: string = '/api/v1/giftpay/apply') => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<ApplyGiftCodeResponse | null>(null);

  const applyCode = async (req: ApplyGiftCodeRequest) => {
    setIsLoading(true);
    setError(null);
    setData(null);

    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(req),
      });

      if (!response.ok) {
        const errorData = (await response.json()) as ErrorResponse;
        throw new Error(errorData.error?.message || 'Failed to apply discount code.');
      }

      const result = (await response.json()) as ApplyGiftCodeResponse;
      setData(result);
      return result;
    } catch (err: any) {
      setError(err.message || 'An unexpected error occurred.');
      throw err;
    } finally {
      setIsLoading(false);
    }
  };

  return {
    applyCode,
    isLoading,
    error,
    data,
    reset: () => {
      setError(null);
      setData(null);
    }
  };
};
