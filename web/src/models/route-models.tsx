type Optional<T> = T | undefined

export interface IRule {
    priority: string
    rule: string
}

export interface IRules {
    rules: Array<IRule> | undefined
}

export class Rule implements IRule {
    priority: string 
    rule: string 
    constructor(){
        this.priority=""
        this.rule=""
    }

    // constructor(arg: Optional<IRule> ) {
    //     if (!arg) {
    //         this.prioriry = undefined
    //         this.rule = undefined
    //     } else {
    //         this.prioriry = arg.prioriry
    //         this.rule = arg.rule
    //     }
    // }
}

export class Rules implements IRules {
    rules: Array<Rule> | undefined

    constructor(arg: Optional<IRules>) {
        this.rules = undefined
    }

    // static FromJSON(json: string) : Optional<Rules> {
    //     try {
    //         let rules: IRules = JSON.parse(json)
    //         if (rules.rules?.length) {
    //             let data = new Rules(rules)
    //         }

    //     } catch (e) {

    //     }
    //     return undefined
    // }
}
//================================================
export interface IRoute {
    index: string 
    route: string 
    tableName:string
}

export interface IRoutes {
    routes: Array<IRoute> | undefined
}

export class Route implements IRoute {
    index: string 
    route: string 
    tableName:string
    constructor(){
        this.index=""
        this.route=""
        this.tableName=""
    }
}

export class Routes implements IRoutes {
    routes: Array<Route> | undefined

    constructor(arg: Optional<IRoutes>) {
        this.routes = undefined
    }
}
//================================================
export interface ITable {
    tableNumber: string
    tableName: string
}

export interface ITables {
    tables: Array<ITable> | undefined
}

export class Table implements ITable {
    tableNumber: string 
    tableName: string 
    constructor(){
        this.tableNumber=""
        this.tableName=""
    }
}

export class Tables implements ITables {
    tables: Array<Table> | undefined

    constructor(arg: Optional<ITables>) {
        this.tables = undefined
    }
}
//================================================
export interface IInterface {
    index: string | undefined
    name: string | undefined
}

export interface IInterfaces {
    interfaces: Array<IInterface> | undefined
}

export class Interface implements IInterface {
    index: string | undefined
    name: string | undefined
}

export class Interfaces implements IInterfaces {
    interfaces: Array<Interface> | undefined

    constructor(arg: Optional<IInterfaces>) {
        this.interfaces = undefined
    }
}
//================================================
export interface IAddress {
    ip: string | undefined
    type: string | undefined
}
export class Address implements IAddress {
    ip: string | undefined
    type: string | undefined
}
export interface IGateway {
    destination: string | undefined
    gateway: string | undefined
}
export class Gateway implements IGateway {
    destination: string | undefined
    gateway: string | undefined
}
//===================================================
export interface IInterfaceAddresses {
    addresses: Array<IAddress> | undefined
}

export class InterfaceAddresses implements IInterfaceAddresses {
    addresses: Array<IAddress> | undefined

    constructor(arg: Optional<IInterfaces>) {
        this.addresses = undefined
    }
}
//======================================================
export interface IInterfaceDetails {
    name:               string | undefined;
    gateways:           Array<IGateway> | undefined;
    normalAddress:      Array<IAddress> | undefined;
    multicastAddress:   Array<IAddress> | undefined;
}
export class InterfaceDetails implements IInterfaceDetails {
    name:               string | undefined;
    gateways:           Array<IGateway> | undefined;
    normalAddress:      Array<IAddress> | undefined;
    multicastAddress:   Array<IAddress> | undefined;
}
//======================================================
export class GeneralRequest {
    sourceIp:       string | undefined;
    tableName:      string | undefined;
    destination:    string | undefined;
    intermediate:   string | undefined;
    interfaceName:  string | undefined;
}

export interface IResponse{
    msg:string | undefined;
}
export class GeneralResponse implements IResponse{
    msg:string | undefined;
}