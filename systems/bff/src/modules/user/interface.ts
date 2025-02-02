import { MetricServiceValueRes, ParsedCookie } from "../../common/types";
import {
    ActivateUserResponse,
    UserResDto,
    AddUserServiceRes,
    ConnectedUserDto,
    DeactivateResponse,
    GetUserDto,
    GetUserResponseDto,
    GetUsersDto,
    OrgUserDto,
    OrgUserResponse,
    OrgUsersResponse,
    UserInputDto,
    GetESimQRCodeInput,
    ESimQRCodeRes,
    UpdateUserServiceInput,
    OrgUserSimDto,
} from "./types";

export interface IUserService {
    getConnectedUsers(cookie: ParsedCookie): Promise<ConnectedUserDto>;
    updateUser(
        userId: string,
        req: UserInputDto,
        cookie: ParsedCookie
    ): Promise<UserResDto>;
    deactivateUser(
        id: string,
        cookie: ParsedCookie
    ): Promise<DeactivateResponse>;
    getUser(userId: string, cookie: ParsedCookie): Promise<GetUserDto>;
    getUsersByOrg(cookie: ParsedCookie): Promise<GetUsersDto[]>;
    addUser(req: UserInputDto, cookie: ParsedCookie): Promise<UserResDto>;
    deleteUser(
        userId: string,
        cookie: ParsedCookie
    ): Promise<ActivateUserResponse>;
    getEsimQRCode(
        data: GetESimQRCodeInput,
        cookie: ParsedCookie
    ): Promise<ESimQRCodeRes>;
    updateUserRoaming(
        data: UpdateUserServiceInput,
        cookie: ParsedCookie
    ): Promise<OrgUserSimDto>;
}

export interface IUserMapper {
    dtoToAddUserDto(res: AddUserServiceRes): UserResDto;
    connectedUsersDtoToDto(res: MetricServiceValueRes[]): ConnectedUserDto;
    dtoToDto(res: GetUserResponseDto): GetUserDto[];
    dtoToUsersDto(req: OrgUsersResponse): GetUsersDto[];
    dtoToUserDto(req: OrgUserResponse): GetUserDto;
    dtoToUserResDto(req: OrgUserDto): UserResDto;
}
