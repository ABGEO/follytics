/* tslint:disable */
/* eslint-disable */
/**
 * Follytics API
 * Follytics API service
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: support@follytics.localhost
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from '../configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, type RequestArgs, BaseAPI, RequiredError, operationServerMap } from '../base';
// @ts-ignore
import type { ResponseHTTPError } from '../model';
// @ts-ignore
import type { ResponseHTTPResponseResponseUser } from '../model';
/**
 * UserApi - axios parameter creator
 * @export
 */
export const UserApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get current authenticated account
         * @summary Get current account
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCurrentUser: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/users/me`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Update the user or create new one if it does not exist
         * @summary Track user login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        trackLogin: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/users/login-events`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * UserApi - functional programming interface
 * @export
 */
export const UserApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = UserApiAxiosParamCreator(configuration)
    return {
        /**
         * Get current authenticated account
         * @summary Get current account
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCurrentUser(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ResponseHTTPResponseResponseUser>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getCurrentUser(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['UserApi.getCurrentUser']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * Update the user or create new one if it does not exist
         * @summary Track user login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async trackLogin(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ResponseHTTPResponseResponseUser>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.trackLogin(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['UserApi.trackLogin']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * UserApi - factory interface
 * @export
 */
export const UserApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = UserApiFp(configuration)
    return {
        /**
         * Get current authenticated account
         * @summary Get current account
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCurrentUser(options?: RawAxiosRequestConfig): AxiosPromise<ResponseHTTPResponseResponseUser> {
            return localVarFp.getCurrentUser(options).then((request) => request(axios, basePath));
        },
        /**
         * Update the user or create new one if it does not exist
         * @summary Track user login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        trackLogin(options?: RawAxiosRequestConfig): AxiosPromise<ResponseHTTPResponseResponseUser> {
            return localVarFp.trackLogin(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * UserApi - object-oriented interface
 * @export
 * @class UserApi
 * @extends {BaseAPI}
 */
export class UserApi extends BaseAPI {
    /**
     * Get current authenticated account
     * @summary Get current account
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof UserApi
     */
    public getCurrentUser(options?: RawAxiosRequestConfig) {
        return UserApiFp(this.configuration).getCurrentUser(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Update the user or create new one if it does not exist
     * @summary Track user login
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof UserApi
     */
    public trackLogin(options?: RawAxiosRequestConfig) {
        return UserApiFp(this.configuration).trackLogin(options).then((request) => request(this.axios, this.basePath));
    }
}

