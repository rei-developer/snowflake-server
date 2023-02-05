import { TypeOrmModuleOptions } from '@nestjs/typeorm';

export const typeORMConfig: TypeOrmModuleOptions = {
  type: 'mariadb',
  host: '127.0.0.1',
  port: 3306,
  username: 'root',
  password: '',
  database: 'snowflake',
  entities: [__dirname + '/../**/*.entity.{js,ts}'],
  // synchronize: true,
};
