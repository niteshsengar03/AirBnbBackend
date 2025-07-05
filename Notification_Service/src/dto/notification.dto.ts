export interface notificationDto{
    to:string;
    subject:string;
    templateId:string;
    params:Record<string,any>;
}