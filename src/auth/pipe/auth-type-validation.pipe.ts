import {
  Injectable,
  PipeTransform,
  ArgumentMetadata,
  BadRequestException,
} from '@nestjs/common';
import { AuthType } from '../auth.model';
import { ExceptionErrorMessage } from '#common/const/exception-error-message.const';

@Injectable()
export class AuthTypeValidationPipe implements PipeTransform {
  private readonly authTypes = Object.keys(AuthType);

  transform(value: any, _: ArgumentMetadata) {
    value = value.toUpperCase();
    if (!this.isAuthTypeValid(value)) {
      throw new BadRequestException(ExceptionErrorMessage.INVALID_PARAMETER);
    }
    return AuthType[value];
  }

  private isAuthTypeValid(value: any) {
    return this.authTypes.indexOf(value) !== -1;
  }
}
