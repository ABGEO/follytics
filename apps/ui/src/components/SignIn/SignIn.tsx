import { Button } from 'antd';
import { LoginOutlined } from '@ant-design/icons';

import { signIn } from '@self/lib/auth';

function SignIn() {
  const logoutAction = async () => {
    'use server';
    await signIn('github', { redirectTo: '/dashboard' });
  };

  return (
    <form action={logoutAction}>
      <Button
        type="primary"
        size="large"
        htmlType="submit"
        icon={<LoginOutlined />}
      >
        Sign In
      </Button>
    </form>
  );
}

export { SignIn };
