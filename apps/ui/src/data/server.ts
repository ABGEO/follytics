type FetchResult<T> = {
  data: T | null;
  error: Error | null;
  isSuccess: boolean;
};

const fetchServerData = async <T, Args extends unknown[]>(
  fetcher: (...args: Args) => Promise<T>,
  ...fetcherArgs: Args
): Promise<FetchResult<T>> => {
  try {
    const data = await fetcher(...fetcherArgs);
    return {
      data,
      error: null,
      isSuccess: true,
    };
  } catch (error) {
    return {
      data: null,
      error: error instanceof Error ? error : new Error(String(error)),
      isSuccess: false,
    };
  }
};

export { fetchServerData, type FetchResult };
