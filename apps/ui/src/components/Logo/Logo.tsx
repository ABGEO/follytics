type LogoProps = {
  width?: string;
  height?: string;
  animate?: boolean;
  animationDuration?: string;
};

function Logo({ width, height, animate, animationDuration = '1s' }: LogoProps) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 200 200"
      fill="currentColor"
      width={width}
      height={height}
    >
      <circle fill="#f5f7fa" cx="100" cy="100" r="90" />
      <path
        fill="#219ebc"
        d="M50,85h71c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z"
      >
        {animate && (
          <animate
            attributeName="d"
            values="M50,85h71c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z;
              M-20,85h71c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z;
              M50,85h71c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z"
            dur={animationDuration}
            repeatCount="indefinite"
          />
        )}
      </path>
      <path
        fill="#8ecae6"
        d="M50,50h96c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z"
      >
        {animate && (
          <animate
            attributeName="d"
            values="M50,50h96c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z;
              M-45,50h96c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z;
              M50,50h96c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z"
            dur={animationDuration}
            repeatCount="indefinite"
          />
        )}
      </path>
      <path
        fill="#023047"
        d="M50,120h31c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4h-31v-30h0Z"
      >
        {animate && (
          <animate
            attributeName="d"
            values="M50,120h31c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4h-31v-30h0Z;
              M20,120h31c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4H50v-30h0Z;
              M50,120h31c2.2,0,4,1.8,4,4v22c0,2.2-1.8,4-4,4h-31v-30h0Z"
            dur={animationDuration}
            repeatCount="indefinite"
          />
        )}
      </path>
    </svg>
  );
}

export { Logo };
