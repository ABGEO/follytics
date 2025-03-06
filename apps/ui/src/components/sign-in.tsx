import { signIn } from 'next-auth/react';

import { Button } from 'antd';

export function SignIn() {
  return (
    <Button onClick={() => signIn('github', { callbackUrl: '/dashboard' })}>
      Sign In
    </Button>
  );
}
