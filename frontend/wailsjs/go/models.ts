export namespace controller {
	
	export class Res {
	    code: number;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Res(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}

}

