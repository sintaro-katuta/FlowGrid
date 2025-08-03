import { ZodError } from 'zod';
import { authSchema } from '$lib/schemas/authSchema.js';
import { fail } from '@sveltejs/kit';

export const actions = {
  login: async ({ request }) => {
    const formData = await request.formData();
    const data = Object.fromEntries(formData);

    try {
      authSchema.pick({ email: true, password: true }).parse(data);
    } catch (err) {
      if (err instanceof ZodError) {
        const { fieldErrors: errors } = err.flatten();
        return fail(400, { errors });
      }
    }

    // const email: string = formData.get('email') as string;
    // const password: string = formData.get('password') as string;

    return { success: true, message: 'ログイン成功' };
  },

  signup: async ({ request }) => {
    const formData = await request.formData();
    const data = Object.fromEntries(formData);

    try {
      authSchema.parse(data);
    } catch (err) {
      if (err instanceof ZodError) {
        const { fieldErrors: errors } = err.flatten();
        return fail(400, { errors });
      }
    }

    // const email: string = formData.get('email') as string;
    // const password: string = formData.get('password') as string;

    return { success: true, message: 'サインアップ成功' };
  }
};
