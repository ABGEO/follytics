import { execSync } from 'child_process';

import { PromiseExecutor, logger } from '@nx/devkit';

import { OpenAPIGeneratorExecutorSchema } from './schema';

const runExecutor: PromiseExecutor<OpenAPIGeneratorExecutorSchema> = async (
  options
) => {
  try {
    const { specFile, generator, output, config } = options;
    const command = `openapi-generator-cli generate -i ${specFile} -g ${generator} -o ${output} -c ${config}`;

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
