type DataPoint = {
  date: Date;
  value: number;
};

type ChartConfiguration = {
  width?: number;
  height?: number;
  backgroundColor?: string;
  lineColor?: string;
  axisColor?: string;
  textColor?: string;
};

type UserDetails = {
  username: string;
  avatar: string;
};

export type { DataPoint, ChartConfiguration, UserDetails };
