import { HeaderType, PaginationDto } from "../../common/types";
import { NetworkDto } from "../network/types";
import { DeactivateResponse } from "../user/types";
import {
    NodeResponseDto,
    NodeResponse,
    NodesResponse,
    AddNodeResponse,
    AddNodeDto,
    OrgNodeResponse,
    OrgNodeResponseDto,
    NodeDetailDto,
    MetricDto,
    OrgMetricValueDto,
    OrgNodeDto,
} from "./types";

export interface INodeService {
    getNodes(req: PaginationDto): Promise<NodesResponse>;
    getNetwork(): Promise<NetworkDto>;
    getNodeDetials(): Promise<NodeDetailDto>;
    getNodesByOrg(
        orgId: string,
        header: HeaderType
    ): Promise<OrgNodeResponseDto>;
    addNode(req: AddNodeDto, header: HeaderType): Promise<AddNodeResponse>;
    updateNode(req: AddNodeDto, header: HeaderType): Promise<OrgNodeDto>;
    deleteNode(id: string): Promise<DeactivateResponse>;
}

export interface INodeMapper {
    dtoToDto(res: NodeResponse): NodeResponseDto;
    dtoToNodesDto(orgId: string, req: OrgNodeResponse): OrgNodeResponseDto;
    dtoToMetricsDto(res: OrgMetricValueDto[]): MetricDto[];
}