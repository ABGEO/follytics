import { Button } from 'antd';

import { signIn } from '@self/lib/auth';

function SignIn() {
  const logoutAction = async () => {
    'use server';
    await signIn('github', { redirectTo: '/dashboard' });
  };

  return (
    <form action={logoutAction}>
      <Button htmlType="submit">Sign In</Button>
    </form>
  );
}

export { SignIn };
