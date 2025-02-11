
package io.fabric8.openshift.api.model.operator.imageregistry.v1;

import java.util.LinkedHashMap;
import java.util.Map;
import javax.annotation.Generated;
import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import io.fabric8.kubernetes.api.builder.Editable;
import io.fabric8.kubernetes.api.model.Container;
import io.fabric8.kubernetes.api.model.IntOrString;
import io.fabric8.kubernetes.api.model.KubernetesResource;
import io.fabric8.kubernetes.api.model.LabelSelector;
import io.fabric8.kubernetes.api.model.LocalObjectReference;
import io.fabric8.kubernetes.api.model.ObjectMeta;
import io.fabric8.kubernetes.api.model.ObjectReference;
import io.fabric8.kubernetes.api.model.PersistentVolumeClaim;
import io.fabric8.kubernetes.api.model.PodTemplateSpec;
import io.fabric8.kubernetes.api.model.ResourceRequirements;
import io.sundr.builder.annotations.Buildable;
import io.sundr.builder.annotations.BuildableReference;
import lombok.EqualsAndHashCode;
import lombok.ToString;
import lombok.experimental.Accessors;

@JsonDeserialize(using = com.fasterxml.jackson.databind.JsonDeserializer.None.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder({
    "networkResourceGroupName",
    "privateEndpointName",
    "subnetName",
    "vnetName"
})
@ToString
@EqualsAndHashCode
@Accessors(prefix = {
    "_",
    ""
})
@Buildable(editableEnabled = false, validationEnabled = false, generateBuilderPackage = false, lazyCollectionInitEnabled = false, builderPackage = "io.fabric8.kubernetes.api.builder", refs = {
    @BuildableReference(ObjectMeta.class),
    @BuildableReference(LabelSelector.class),
    @BuildableReference(Container.class),
    @BuildableReference(PodTemplateSpec.class),
    @BuildableReference(ResourceRequirements.class),
    @BuildableReference(IntOrString.class),
    @BuildableReference(ObjectReference.class),
    @BuildableReference(LocalObjectReference.class),
    @BuildableReference(PersistentVolumeClaim.class)
})
@Generated("jsonschema2pojo")
public class AzureNetworkAccessInternal implements Editable<AzureNetworkAccessInternalBuilder> , KubernetesResource
{

    @JsonProperty("networkResourceGroupName")
    private String networkResourceGroupName;
    @JsonProperty("privateEndpointName")
    private String privateEndpointName;
    @JsonProperty("subnetName")
    private String subnetName;
    @JsonProperty("vnetName")
    private String vnetName;
    @JsonIgnore
    private Map<String, Object> additionalProperties = new LinkedHashMap<String, Object>();

    /**
     * No args constructor for use in serialization
     * 
     */
    public AzureNetworkAccessInternal() {
    }

    public AzureNetworkAccessInternal(String networkResourceGroupName, String privateEndpointName, String subnetName, String vnetName) {
        super();
        this.networkResourceGroupName = networkResourceGroupName;
        this.privateEndpointName = privateEndpointName;
        this.subnetName = subnetName;
        this.vnetName = vnetName;
    }

    @JsonProperty("networkResourceGroupName")
    public String getNetworkResourceGroupName() {
        return networkResourceGroupName;
    }

    @JsonProperty("networkResourceGroupName")
    public void setNetworkResourceGroupName(String networkResourceGroupName) {
        this.networkResourceGroupName = networkResourceGroupName;
    }

    @JsonProperty("privateEndpointName")
    public String getPrivateEndpointName() {
        return privateEndpointName;
    }

    @JsonProperty("privateEndpointName")
    public void setPrivateEndpointName(String privateEndpointName) {
        this.privateEndpointName = privateEndpointName;
    }

    @JsonProperty("subnetName")
    public String getSubnetName() {
        return subnetName;
    }

    @JsonProperty("subnetName")
    public void setSubnetName(String subnetName) {
        this.subnetName = subnetName;
    }

    @JsonProperty("vnetName")
    public String getVnetName() {
        return vnetName;
    }

    @JsonProperty("vnetName")
    public void setVnetName(String vnetName) {
        this.vnetName = vnetName;
    }

    @JsonIgnore
    public AzureNetworkAccessInternalBuilder edit() {
        return new AzureNetworkAccessInternalBuilder(this);
    }

    @JsonIgnore
    public AzureNetworkAccessInternalBuilder toBuilder() {
        return edit();
    }

    @JsonAnyGetter
    public Map<String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    @JsonAnySetter
    public void setAdditionalProperty(String name, Object value) {
        this.additionalProperties.put(name, value);
    }

    public void setAdditionalProperties(Map<String, Object> additionalProperties) {
        this.additionalProperties = additionalProperties;
    }

}
