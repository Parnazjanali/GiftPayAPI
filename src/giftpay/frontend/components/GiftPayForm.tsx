import React, { useState, FormEvent } from 'react';
import { useGiftPay } from '../hooks/useGiftPay';

interface GiftPayFormProps {
  userId: string;
  onSuccess?: (newBalance?: number) => void;
}

export const GiftPayForm: React.FC<GiftPayFormProps> = ({ userId, onSuccess }) => {
  const [code, setCode] = useState('');
  const [shake, setShake] = useState(false);
  const { applyCode, isLoading, error, data, reset } = useGiftPay();

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    if (!code.trim()) {
      triggerShake();
      return;
    }

    try {
      const result = await applyCode({ code: code.trim(), userId });
      if (onSuccess) {
        onSuccess(result.newBalance);
      }
    } catch (err) {
      triggerShake();
    }
  };

  const triggerShake = () => {
    setShake(true);
    setTimeout(() => setShake(false), 500);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCode(e.target.value.toUpperCase());
    if (error || data) {
      reset();
    }
  };

  return (
    <div className="w-full max-w-md mx-auto p-6 bg-white/80 backdrop-blur-md border border-gray-100 rounded-2xl shadow-xl shadow-amber-900/5">
      <div className="text-center mb-6">
        <h2 className="text-2xl font-semibold text-gray-800">Redeem Gift Code</h2>
        <p className="text-sm text-gray-500 mt-1">Enter your code below to add to your balance.</p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div className={`relative transition-transform duration-300 ${shake ? 'animate-[shake_0.5s_ease-in-out]' : ''}`}>
          <label htmlFor="gift-code" className="sr-only">
            Gift Code
          </label>
          <input
            id="gift-code"
            type="text"
            value={code}
            onChange={handleInputChange}
            disabled={isLoading || !!data}
            placeholder="e.g. SUMMER50"
            className={`w-full px-4 py-3 rounded-xl border ${
              error ? 'border-red-300 focus:ring-red-500/20' : 'border-gray-200 focus:ring-amber-500/20 focus:border-amber-500'
            } bg-white/50 focus:bg-white text-gray-800 font-medium tracking-wide placeholder:text-gray-400 placeholder:font-normal outline-none transition-all`}
            aria-invalid={!!error}
            aria-describedby={error ? "code-error" : undefined}
          />
        </div>

        {error && (
          <p id="code-error" className="text-sm text-red-500 font-medium px-1 animate-[fadeIn_0.3s_ease-in]">
            {error}
          </p>
        )}

        {data && (
          <div className="p-4 bg-green-50 border border-green-100 rounded-xl animate-[slideDown_0.3s_ease-out]">
            <p className="text-green-800 font-medium flex items-center gap-2">
              <svg className="w-5 h-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 13l4 4L19 7" />
              </svg>
              Success! Added {data.appliedValue} to your balance.
            </p>
            {data.newBalance !== undefined && (
              <p className="text-sm text-green-600 mt-1 ml-7">
                New Balance: <span className="font-bold">{data.newBalance}</span>
              </p>
            )}
          </div>
        )}

        {!data && (
          <button
            type="submit"
            disabled={isLoading || !code.trim()}
            className="w-full relative py-3 px-4 bg-gradient-to-r from-amber-500 to-amber-600 hover:from-amber-600 hover:to-amber-700 disabled:from-gray-300 disabled:to-gray-400 text-white font-semibold rounded-xl shadow-lg shadow-amber-500/30 hover:shadow-amber-600/40 disabled:shadow-none transition-all duration-200 active:scale-[0.98] disabled:active:scale-100 focus:outline-none focus:ring-2 focus:ring-amber-500/50 focus:ring-offset-2"
          >
            {isLoading ? (
              <span className="flex items-center justify-center gap-2">
                <svg className="animate-spin h-5 w-5 text-white/90" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                  <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Applying...
              </span>
            ) : (
              'Apply Code'
            )}
          </button>
        )}
      </form>

      <style dangerouslySetInnerHTML={{__html: `
        @keyframes shake {
          0%, 100% { transform: translateX(0); }
          25% { transform: translateX(-4px); }
          50% { transform: translateX(4px); }
          75% { transform: translateX(-4px); }
        }
        @keyframes slideDown {
          from { opacity: 0; transform: translateY(-10px); }
          to { opacity: 1; transform: translateY(0); }
        }
        @keyframes fadeIn {
          from { opacity: 0; }
          to { opacity: 1; }
        }
      `}} />
    </div>
  );
};
