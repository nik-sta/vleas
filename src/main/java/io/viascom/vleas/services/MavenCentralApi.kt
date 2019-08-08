package io.viascom.vleas.services

import ch.viascom.groundwork.foxhttp.FoxHttpResponse
import ch.viascom.groundwork.foxhttp.builder.FoxHttpClientBuilder
import ch.viascom.groundwork.foxhttp.builder.FoxHttpRequestBuilder
import ch.viascom.groundwork.foxhttp.parser.GsonParser
import ch.viascom.groundwork.foxhttp.type.RequestType

object MavenCentralApi {
    fun retrieveDependency(group: String, name: String): FoxHttpResponse {
        return retrieveDependencies(group,name,1)
    }

    fun retrieveDependencies(group: String, name: String, results: Int): FoxHttpResponse {
        val client = FoxHttpClientBuilder(GsonParser()).build()
        return FoxHttpRequestBuilder("http://search.maven.org/solrsearch/select?q=g:%22$group%22+AND+a:%22$name%22&rows=$results&core=gav",RequestType.GET, client)
            .build()
            .execute()
    }
}