import { z } from 'zod';

export const authSchema = z.object({
  email: z.email({ message: '有効なEメールアドレスを入力してください' }),
  password: z.string().min(8, { message: 'パスワードは8文字以上で入力してください' })
});

export type authSchemaType = typeof authSchema;
