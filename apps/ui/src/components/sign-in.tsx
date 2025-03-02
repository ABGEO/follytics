import { signIn } from 'next-auth/react';

import { Button } from '@self/components/ui/button';

export function SignIn() {
  return <Button onClick={() => signIn('github')}>Sign In</Button>;
}
