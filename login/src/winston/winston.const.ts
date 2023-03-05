import { appConfig } from '#configs/app.config';

const {
  log: { dir, maxFiles, zippedArchive },
} = appConfig;

export const dailyOptions = (level: string) => {
  return {
    level,
    datePattern: 'YYYY-MM-DD',
    dirname: `${dir}/${level}`,
    filename: `%DATE%.${level}.log`,
    maxFiles,
    zippedArchive,
  };
};
