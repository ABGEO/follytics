'use client';

import Image from 'next/image';
import { useSession } from 'next-auth/react';

import { Avatar, AvatarFallback, AvatarImage } from '@self/components/ui/avatar';
import { SignIn } from '@self/components/sign-in';
import { SignOut } from '@self/components/sign-out';

export default function Home() {
  const { data: session, status } = useSession();

  return (
    <main className="flex flex-col gap-8 row-start-2 justify-center h-screen justify-items-center items-center">
      <div className="flex gap-4 items-center flex-col ">
        {status !== 'authenticated' && <SignIn />}

        {status === 'authenticated' && (
          <>
            <Avatar className="w-14 h-14">
              <AvatarImage asChild src={session.user?.image as string}>
                <Image
                  src={session.user?.image as string}
                  alt={session.user?.name as string}
                  width={128}
                  height={128}
                />
              </AvatarImage>
              <AvatarFallback>{session.user?.name}</AvatarFallback>
            </Avatar>

            <h1 className="text-xl tracking-tight">
              Hello, {session.user?.name}!
            </h1>
            <SignOut />
          </>
        )}
      </div>
    </main>
  );
}
