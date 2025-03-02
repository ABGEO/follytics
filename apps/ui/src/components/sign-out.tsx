import { Button } from '@self/components/ui/button';
import { signOut } from 'next-auth/react';

export function SignOut() {
  return <Button onClick={() => signOut()}>Sign Out</Button>;
}
