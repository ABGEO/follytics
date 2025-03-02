import { execSync } from 'child_process';

import { PromiseExecutor, logger } from '@nx/devkit';

import { SwagExecutorSchema } from './schema';

const runExecutor: PromiseExecutor<SwagExecutorSchema> = async (options) => {
  try {
    const { outputTypes, output, dirs, generalInfo } = options;

    const command = `swag init --ot ${outputTypes.join(
      ','
    )} -o ${output} -d ${dirs.join(',')} -g ${generalInfo}`;

    logger.info(`Executing command: ${command}`);

    execSync(command, {
      stdio: [0, 1, 2],
    });

    return { success: true };
  } catch (error) {
    logger.error(error);

    return { success: false };
  }
};

export default runExecutor;
